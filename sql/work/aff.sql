------------------------------------------------------------------------------------------------------------------------
--  DATABASE TABLES SIZES
------------------------------------------------------------------------------------------------------------------------
/*
SELECT
    table_name AS table_name,
    engine,
    ROUND(data_length/1024/1024,2) AS total_size_mb,
    table_rows
FROM
    information_schema.tables
WHERE
    table_schema=DATABASE();
*/
SELECT table_name, ROUND((data_length + data_free + index_length) / (1 << 30), 2) AS size 
FROM information_schema.tables WHERE table_schema = 'aff_easydate_biz' 
GROUP BY table_name ORDER BY size DESC LIMIT 30;

------------------------------------------------------------------------------------------------------------------------
--  TABLES OVERFLOW
------------------------------------------------------------------------------------------------------------------------
select 1<<8, 1<<32, MAX(rule_queueid), (1 << 32) - MAX(rule_queueid) from mail_rule_queue;
+------+------------+-------------------+--------------------------------+
| 1<<8 | 1<<32      | MAX(rule_queueid) | (1 << 32)  - MAX(rule_queueid) |
+------+------------+-------------------+--------------------------------+
|  256 | 4294967296 |        2154452210 |                     2140515086 |
+------+------------+-------------------+--------------------------------+
-- [21:51:44] Evgeniy Petrov: 1 << 31 смотри
-- [21:51:52] Evgeniy Petrov: 1 << 32 - это беззнаковый

------------------------------------------------------------------------------------------------------------------------
-- REPORTS
------------------------------------------------------------------------------------------------------------------------
-- Shows sites ratings by count of transactions per period
select left(tsa.productid,locate('_',tsa.productid)-1) site, count(tsa.transid)cnt
from transactions_sales_alcuda tsa
where tsa.dateinserted between date(now()-interval 15 day) and date(now()-interval 1 day)
group by site
order by cnt
;

------------------------------------------------------------------------------------------------------------------------
-- TRANSACTIONS
------------------------------------------------------------------------------------------------------------------------
-- Shows diff between transactions and quick report.
set @date1 = date(now()-interval 2 day);
set @date2 = date(now()-interval 1 day);
set @date1 = '2013-05-01 00:00:00';
set @date2 = '2013-05-30 23:59:59';
-- part #1
select transtype,count(transtype) from transactions_sales_alcuda where dateinserted between @date1 and @date2 and shaved=1 group by transtype;
select sum(reg_shaved),sum(leads_shaved),sum(sales_shaved) from users_sales_detailed where date between @date1 and @date2 and (reg_shaved>0 or leads_shaved>0 or sales_shaved>0);
-- part #2
select affiliateid,count(affiliateid)cnt from transactions_sales_alcuda where dateinserted between @date1 and @date2 and shaved=1 and transtype in (4) group by affiliateid order by affiliateid;
select t1.userid,count(t0.user_data_id) cnt,t0.user_data_id from users_sales_detailed t0 join users_data t1 on t0.user_data_id=t1.id and t1.accountid='alcuda' where t0.date between @date1 and @date2 and (t0.sales_shaved>0) group by t0.user_data_id order by t1.userid;
select distinct t0.user_data_id from users_sales_detailed t0 join users_data t1 on t0.user_data_id=t1.id and t1.accountid='alcuda' where t0.date between @date1 and @date2 ;



-- 
select u.userid,b.userid,b.backurl,round(sum(s.approved),2) approved_amount
from users_data u join wd_g_backtraffic b on u.userid=b.userid join users_sales_detailed s on u.id=s.user_data_id
where accountid='alcuda' and b.enabled=1 and b.backurl>'' and b.backurl not like '%a_aid%' and s.date>'2013' and s.approved>0
group by u.userid
order by approved_amount;


--  Shows transactions with 1 order id that tracked more than 1.
select affiliateid,orderid,productid,count(orderid) cnt,group_concat(transid) -- ,transtype,countrycode
from transactions_sales_alcuda 
where dateinserted between '2013-01-25 00:00:00' and '2013-01-31 23:59:59' and hidden=0 -- and rstatus in (1, 2)
group by orderid,productid -- ,countrycode
having cnt>1 
order by productid
;
-- find transids for clear db
select transid from transactions_sales_alcuda where hidden=0 and rstatus in (1, 2) and orderid=232823613;
-- clear db
update transactions_sales_alcuda set rstatus=3, hidden=1 where transid in(
3963326,3962001,3961178,3961265,3961267,3961268,3961266,3961273,3961274,3961532,3962033,3962172,3962187,3962282,3962330,3962526,3962529,3962528,3963472,3963002,3963000,3962993,3961999,3962843,3961324,3962270,3961224,3961240,3961228,3961476,3961696,3961789,3961915,3962182,3962335,3962329,3962342,3962407,3962446,3962829,3963010,3963015,3963184,3963237,3963271,3963336,3963334,3963435
);


-- Users with clicks count smaller than leads count.
select t1.user_data_id,ud. userid,name,surname, date(t1.date) d,t1.clicks,t1.leads
from users_trans_stats t1 join users_data ud on t1.user_data_id=ud.id
where t1.date between (now()-interval 2 day) and (now()-interval 1 day) and t1.clicks<t1.leads group by t1.user_data_id,t1.date
order by leads desc;
-- Show bad records
select left(str,locate('&dte=',str)) s, count(*) cnt from _click_debug where str like '%a_bid=&%' group by s order by cnt desc;
-- Show clicks input errors with empty a_bid.
select message,count(message) cnt from transactions_errors where date>date(now()-interval 7 day) and message like 'Param a_bid is missing. a_aid%' group by message order by cnt desc;


-- этот запрос показывает сбриваются ли реферальные транзакции (т.е. нормальные транзакции у реферального афилейта)
-- учитывается id реферального афилейта = transactions_sales_alcuda.parent_id
set @v_id     = 76656;
set @v_date_1 = '2012-06-1 00:00:00';
set @v_date_2 = '2012-06-30 23:59:59';
select 
    t.transid,
    t.affiliateid,
    u.id,
    count(t.countrycode)    as `amount countries`,
    count(c.tracking_shave) as `amount shave 1`,
    count(ts_def.percent)   as `amount shave 2`,
    count(ts_uc.percent)    as `amount shave 3`
from transactions_sales_alcuda t
join users_data u on t.affiliateid = u.userid and u.accountid = 'alcuda'
left join countries c on t.countrycode = c.code
left join tracking_shave ts_def on u.id = ts_def.for_id and ts_def.country is null
left join tracking_shave ts_uc on u.id = ts_uc.for_id and ts_uc.country = t.countrycode
where t.parent_id = @v_id
and t.dateinserted between @v_date_1 and @v_date_2
and t.shaved = 1
group by t.affiliateid
;

-- TRANSACTIONS REPORT
SELECT
    t1.transid, t1.commission, t1.totalcost, t1.productid, t1.dateinserted, t1.transtype, t1.affiliateid, t1.rstatus, t1.payoutstatus, t1.ip, t1.countrycode, t1.campcategoryid, t1.datepayout, t1.count, t1.decline_reason, t1.date_status, t1.subnet_id, t1.hidden, t1.currency_rate, t1.shaved, t2.username, t2.name, t2.surname, t2.weburl, t2.accountid, t2.trusted, t3.userid 
FROM 
    campaigncategories t0 
    INNER JOIN transactions_sales_alcuda t1 ON t0.campcategoryid = t1.campcategoryid 
    INNER JOIN users_data t2 ON t1.affiliateid = t2.userid 
    INNER JOIN wd_g_users t3 ON t2.userid = t3.userid 
WHERE 
    t0.campaignid IN ('0158831f', '27748a5d', '8eb302b7', 'b06deb8e', 'e0fbc8a1', '337f53a7', 'e46f6494', '9b80d209', '985803a6') 
    AND t1.countrycode = 'US' 
    AND (t1.dateinserted BETWEEN '2012-7-1 00:00:00' AND '2012-7-31 23:59:59') 
    AND t1.affiliateid = '2d2e8fc0' 
    AND t1.rstatus IN (1, 2, 3) 
    AND t1.hidden IN (0, 1) 
    AND t1.campaignid = 'b06deb8e' 
    AND t1.transtype IN (1, 2, 4, 8, 16) 
    AND t2.accountid = 'alcuda' 
    AND t3.deleted = 0 
    AND t3.rstatus IN (2, 1) 
ORDER BY dateinserted desc 
LIMIT 0, 20
;

-- Показывает сколько зашло транзакции (кроме кликов) за день по часам
select hour(date) as hour, count(id) 
from _sale_debug 
where date > '2012-08-28' 
group by hour 
order by hour;
-- Показывает сколько зашло транзакции (кроме кликов) per days
set SQL_BIG_SELECTS=1;
select date(date)d, count(id)
from _sale_debug 
where date between date(now()-interval 0 day) and date(now()+interval 1 day)
group by d 
;
------------------------------------------------------------------------------------------------------------------------

select count(*) from (
    select count(str) as cnt,str from (
        select 
        concat( if(`dateinserted` is null,'',`dateinserted`),
            if(`bannerid` is null,'',`bannerid`),
            if(`affiliateid` is null,'',`affiliateid`),
            if(`ip` is null,'',`ip`),
            if(`countrycode` is null,'',`countrycode`),
            if(`count` is null,'',`count`),
            if(`datacenter` is null,'',`datacenter`),
            if(`subnet_id` is null,'',`subnet_id`),
            if(`user_data_id` is null,'',`user_data_id`),
            if(`siteid` is null,'',`siteid`),
            if(`unique` is null,'',`unique`)
        ) as str
        from transactions_clicks_alcuda where dateinserted>'2012-09-20 0' and dateinserted<'2012-09-20 23'
    ) as tbl
    group by str
    having cnt > 1
    -- limit 11
) as res
;


------------------------------------------------------------------------------------------------------------------------
-- TRANSACTIONS DEBUG
------------------------------------------------------------------------------------------------------------------------
-- Displays transactions available at sale_debug but not allowed at transactions.
set @date1 = '2013-11-15';
set @date2 = '2013-11-25';
select
    -- count(d.orderId)cnt,
    -- d.orderId,notPay, ts.orderid,affiliateid,decline_reason,hidden,shaved,countrycode,productid,rstatus,transtype
    d.day,count(d.day)
    -- ,group_concat(d.orderId)
    -- str
from (
    select
        date(sd.date)day,
        right(left(str, locate('OrderID=', str)+16), 9)orderId,
        if (str not like '%comP%', 1, 0) notPay,
        te.id,
        str
    from _sale_debug sd
    left join transactions_errors te on sd.id=te.debug_id
    where sd.date between @date1 and @date2 and sd.str like '%OrderID%'
    and te.id is null
    -- limit 11;
)d
-- left join  transactions_sales_alcuda ts on ts.orderid = d.OrderID
left join  transactions_sales_enedina ts on ts.orderid = d.OrderID
where ts.orderid is null
and d.notPay = 1
and d.orderId is not null
group by d.day
-- group by d.orderId
-- order by ts.orderid, d.notPay
-- limit 11
;



/*
_sale_debug_decline
_sale_debug_paytrack
_sale_debug
*/
-- sale debug Кол-во полученных транзакции по часам
select hour(date) hour,count(hour(date)) from _sale_debug where date between date(now()-interval 1 day) and date(now()) group by hour;
-- sale debug Кол-во полученных кликов по часам
select HOUR(date) as hour,count(date) from _click_debug where date > CURRENT_DATE group by hour;
-- кол-во полученных транзакции в разрезе по аффилейтам
select RIGHT(LEFT(str, LOCATE('a_aid=', str)+13), 8) as a_aid, count(str) 
from _sale_debug where date > '2012-09-11 01:00:00' -- BETWEEN '2012-09-11 01:00:00' AND '2012-09-11 01:00:00'
group by a_aid order by a_aid ;
-- выгребем ордеры полученных платежей
select right(left(str, locate('orderid=', str)+15), 8) as ordrs 
from _sale_debug where str like '%_comp%' and date between date(now()-interval 1 day) and date(now())
;
-- PAYMENT COUNT BY SITES
select site, count(site) cnt, group_concat(ordr) from (
    select 
        substr(str, locate('ProductID=', str)+10, locate('&', str, locate('ProductID=', str))-16) site,
        right(left(str, locate('orderid=', str)+15), 8) as ordr
    from _sale_debug where str like '%_comp%internal%' and date between date(now()-interval 2 day) and date(now()-interval 1 day)
) t group by site /*with rollup*/ order by cnt desc
;

------------------------------------------------------------------------------------------------------------------------
-- USERS
------------------------------------------------------------------------------------------------------------------------
-- All about AFF.
select
    ud.username mail, ud.name, ud.surname, ud.id, ud.userid, ud.accountid, ud.signup_ip, ud.last_login
    -- ,group_concat(concat(lh.date_inserted,'--',lh.ip) separator '; ')
    , if(wgu.source_type=0,'dirct', if(wgu.source_type=1,'cpa', 'else') ) source
from users_data ud
left join wd_g_users wgu on ud.userid=wgu.userid
-- left join login_history lh on ud.id=lh.user_data_id and lh.is_logged=1
where ud.userid in ('6c553139','be0be0ed','dea02327', '7ef24092');

-- Показывает когда менеджеры логинилсь последний раз в админку
SELECT u.userid, ud.last_login, username, name, surname, accountid
FROM wd_g_users AS u 
LEFT JOIN users_data AS ud ON u.userid = ud.userid 
WHERE u.rtype=3 AND ud.accountid = 'alcuda'
ORDER BY last_login
;

set @v_aid = '769dfde7';
SELECT q.id,q.userid,q.accountid,q.refid,q.username,q.rpassword,q.name,q.surname,w.alhash,w.userid,w.role_id 
FROM users_data AS q 
LEFT JOIN wd_g_users AS w 
ON q.userid=w.userid 
WHERE q.userid = @v_aid;

set @v_aid = '127690761';
SELECT q.id,q.userid,q.accountid,q.refid,q.username,q.rpassword,q.name,q.surname,w.alhash,w.userid,w.role_id 
FROM users_data AS q 
LEFT JOIN wd_g_users AS w 
ON q.userid=w.userid 
WHERE q.id=@v_aid;


SELECT q.id,q.userid,q.accountid,q.refid,q.username,q.rpassword,q.name,q.surname,w.alhash,w.userid,w.role_id,w.source_type
FROM users_data AS q 
LEFT JOIN wd_g_users AS w 
ON q.userid=w.userid 
WHERE q.id='80894';

SELECT q.id,q.userid,q.accountid,q.refid,q.username,q.rpassword,q.name,q.surname,w.alhash,w.userid,w.role_id,w.source_type
FROM users_data AS q 
LEFT JOIN wd_g_users AS w 
ON q.userid=w.userid 
WHERE q.userid='admin_29';

SELECT q.id,q.userid,q.accountid,q.refid,q.username,q.rpassword,q.name,q.surname,w.alhash,w.userid,w.role_id 
FROM users_data AS q 
LEFT JOIN wd_g_users AS w 
ON q.userid=w.userid 
WHERE q.username='vladimitk@ufins.com';

SELECT q.id,q.userid,q.accountid,q.refid,q.username,q.rpassword,q.name,q.surname,w.alhash,w.userid,w.role_id 
FROM users_data AS q 
LEFT JOIN wd_g_users AS w 
ON q.userid=w.userid 
WHERE w.alhash='341de58cb2346e60a561d074ac5671df';

SELECT count(*)
-- q.id,q.userid,q.accountid,q.refid,q.username,q.rpassword,q.name,q.surname,w.alhash,w.userid,w.role_id,w.source_type
FROM users_data AS q 
LEFT JOIN wd_g_users AS w 
ON q.userid=w.userid 
WHERE w.source_type='1';


------------------------------------------------------------------------------------------------------------------------
-- Shows count of registred affiliates at aff.sys and count of who become active or nor.
-- When aff receive first lead faster that 180 days (6 months) it will has active status.
set sql_big_selects=1;
set @date1 = '2011-01-01';
set @date2 = '2014-01-01';
-- explain
select
    left(dateinserted,7)month,
    sum(active)active,
    sum(not_active)not_active,
    sum(passive)still_passive,
    count(dateinserted)total
from (
    select
        userid,
        dateinserted,
        date_first_lead,
        if(passive>0 or date_first_lead is null, 1, 0)passive,
        if(date_first_lead is not null and dateinserted is not null and datediff(date_first_lead,dateinserted)<=180, 1, 0)active,
        if(date_first_lead is not null and dateinserted is not null and datediff(date_first_lead,dateinserted)>180 , 1, 0)not_active
    from (
        select
            ud.id,
            ud.userid,
            wgu.dateinserted,
            ud.start_working_date,
            min(usd.date)date_first_lead,
            if(ud.start_working_date is null, 1, 0)passive
        from users_data ud
        join wd_g_users wgu on ud.userid=wgu.userid
        left join users_sales_detailed usd on ud.start_working_date is not null and ud.id=usd.user_data_id and usd.leads>0
        where wgu.dateinserted between @date1 and @date2 and ud.accountid='alcuda' and wgu.rtype=4
        group by ud.userid
    )t
)t_agr
group by month
;
-- Not aggregated data, allows to see real numbers per short period.
set sql_big_selects=1;
set @date1 = '2013-09-01';
set @date2 = '2013-10-01';
-- explain
select
    userid,
    dateinserted,
    date_first_lead,
    datediff(date_first_lead,dateinserted)diff_days,
    if(passive>0 or date_first_lead is null, 1, 0)passive,
    if(date_first_lead is not null and dateinserted is not null and datediff(date_first_lead,dateinserted)<=180, 1, 0)active,
    if(date_first_lead is not null and dateinserted is not null and datediff(date_first_lead,dateinserted)>180 , 1, 0)not_active
from (
    select
        ud.id,
        ud.userid,
        wgu.dateinserted,
        ud.start_working_date,
        min(usd.date)date_first_lead,
        if(ud.start_working_date is null, 1, 0)passive
    from users_data ud
    join wd_g_users wgu on ud.userid=wgu.userid
    left join users_sales_detailed usd on ud.start_working_date is not null and ud.id=usd.user_data_id and usd.leads>0
    where wgu.dateinserted between @date1 and @date2 and ud.accountid='alcuda' and wgu.rtype=4
    group by ud.userid
)t
having date_first_lead is not null and dateinserted is not null
;
------------------------------------------------------------------------------------------------------------------------
-- Shows aff came from url
-- set sql_big_selects=1;
-- set @date1 = '2013-01-01';
-- set @date2 = '2013-09-01';
-- select
--     host,
--     currency,
--     count(aff)aff_cnt,
--     round(sum(pending),2)pending,
--     round(sum(approved),2)approved,
--     round(sum(paid),2)paid,
--     group_concat(aff)
-- from (
--     select
--         ud.userid aff, usd.currency, sum(usd.pending)pending, sum(usd.approved)approved, sum(usd.paid)paid,
--         left(
--             right(ae.ref, length(ae.ref)-locate('://',ae.ref)-2),
--             locate('/', right(ae.ref, length(ae.ref)-locate('://',ae.ref)-2) )-1
--         ) host
--     from users_sales_detailed usd 
--     join users_data ud on usd.user_data_id = ud.id and ud.accountid='alcuda'
--     join wd_g_users wgu using(userid)
--     left join aff_entry ae on ud.userid = ae.affiliateid
--     where wgu.dateinserted between @date1 and @date2 and usd.date between @date1 and @date2
--     group by user_data_id,currency
-- )t
-- group by host,currency
-- order by aff_cnt desc, host
-- ;
set sql_big_selects=1;
set @date1 = '2013-01-01';
set @date2 = '2013-09-01';
select
    host,
    currency,
    count(aff)aff_cnt,
    round(sum(pending),2)pending,
    round(sum(approved),2)approved,
    round(sum(paid),2)paid,
    group_concat(aff)
from (
    select
        ud.userid aff, usd.currency, sum(usd.pending)pending, sum(usd.approved)approved, sum(usd.paid)paid,
        left(
            right(ae.ref, length(ae.ref)-locate('://',ae.ref)-2),
            locate('/', right(ae.ref, length(ae.ref)-locate('://',ae.ref)-2) )-1
        ) host
    from aff_entry ae
    join wd_g_users wgu on ae.affiliateid = wgu.userid and wgu.rtype=4
    join users_data ud on ud.userid = wgu.userid and ud.accountid='alcuda'
    left join users_sales_detailed usd on usd.user_data_id = ud.id
    where wgu.dateinserted between @date1 and @date2
    group by aff,currency
)t
group by host,currency
order by aff_cnt desc, host
;
------------------------------------------------------------------------------------------------------------------------
-- MAILINGS
------------------------------------------------------------------------------------------------------------------------
set @v_mailing_id = 79; -- AM
set @v_mailing_id = 72; -- GA
set @v_mailing_history_id = 149;
set @v_date = '2013-01-01 00:00:00';

insert into mailings_queue set user_data_id=79838, mailing_id=@v_mailing_id, date_inserted=now(), mailing_history_id=@v_mailing_history_id; -- vladimirk@ufins.com
insert into mailings_queue set user_data_id=80603, mailing_id=@v_mailing_id, date_inserted=now(), mailing_history_id=@v_mailing_history_id; -- nadya@upforitnetworks.com
insert into mailings_queue set user_data_id=78636, mailing_id=@v_mailing_id, date_inserted=now(), mailing_history_id=@v_mailing_history_id; -- anna@ufins.com
insert into mailings_queue select null,id,@v_mailing_id,now(),@v_mailing_history_id from users_data where username like '%codenamek2010%';

-- Ставим всех алькудовских пользователей на емейл рассылку с id # @v_mailing_id
INSERT INTO
    mailings_queue 
SELECT
    null as id,
    ud.id as user_data_id,
    @v_mailing_id as mailing_id,
    now() as date_inserted,
    @v_mailing_history_id as mailing_history_id
FROM 
    wd_g_users u
    INNER JOIN users_data ud ON ud.userid = u.userid
WHERE
    ud.accountid = 'enedina' -- alcuda
    AND u.rtype = 4
    AND ud.subscriber = 1
    AND ud.last_login > @v_date
;
-- Кол-во всех пользователей алькуды, которые попадут в рассылку
SELECT count(ud.id)
FROM 
    wd_g_users u
    INNER JOIN users_data ud ON ud.userid = u.userid
WHERE
    ud.accountid = 'enedina' -- alcuda
    AND u.rtype = 4
    AND ud.subscriber = 1
    AND ud.last_login > @v_date
;
/* -- Если какие-то пользователи не попали в рассылку
INSERT INTO
    mailings_queue 
SELECT
    null as id,
    ud.id as user_data_id,
    @v_mailing_id as mailing_id,
    null as date_inserted,
    null as mailing_history_id
FROM 
    wd_g_users u
    INNER JOIN users_data ud ON ud.userid = u.userid
WHERE
    ud.accountid = 'alcuda'
    AND u.rtype = 4
    AND ud.subscriber = 1
    AND ud.last_login > @v_date
    AND ud.id NOT IN (SELECT user_data_id FROM mailings_queue)
;
*/
-- Проверяем кол-во вставленых в очередь аффилейтов
SELECT count(*) FROM mailings_queue;
-- Проверяем номер рассылки
SELECT distinct mailing_id FROM mailings_queue;
SELECT count(distinct mailing_id) FROM mailings_queue;
-- Проверим не дублируются ли пользователи в очереди на рассылку
SELECT user_data_id, count(user_data_id) as ud_id FROM mailings_queue group by user_data_id having ud_id > 1;
DELETE FROM mailings_queue WHERE id>1;

------------------------------------------------------------------------------------------------------------------------
--  Managers activity
set @date1 = date(now()-interval 60 day);
set @date2 = date(now()-interval 2 day);

select
    l2.action,la.name,group_concat(distinct ud.surname),count(l2.action)
from log_manager2 l2 join log_actions la on l2.action=la.id join users_data ud on l2.manager_id=ud.id
where l2.date between @date1 and @date2
group by l2.action;

select * from log_manager2 where date between @date1 and @date2 and action in (15);

select l2.action,la.name,group_concat(distinct ud.surname),count(l2.action)
from log_manager2 l2 join log_actions la on l2.action=la.id join users_data ud on l2.manager_id=ud.id
where l2.date between date(now()-interval 1 day) and date(now()) and (l2.debug like '%45988315%' or l2.message like '%45988315%');



------------------------------------------------------------------------------------------------------------------------
-- BL_Report->getTrans
explain
select * from (
    ( SELECT t0.transid as transid, t0.transtype as transtype, t0.dateinserted as dateinserted, t0.countrycode as countrycode, t0.subnet_id as subnet_id, t0.totalcost as totalcost, t0.bannerid as bannerid, t0.rstatus as rstatus, t0.ip as ip, t0.productid as gender, t0.orderid as orderid, t0.payoutstatus as payoutstatus, t0.commission as commission, t0.currency_rate as currency_rate, t0.parent_id as parent_id, dr.reason as decline_reason, REPLACE( st.name,".com","") as site, ct.name as countryname
    FROM transactions_sales_alcuda t0 
    LEFT JOIN decline_reasons dr ON dr.id = t0.decline_reason 
    LEFT JOIN wd_pa_campaigns cm ON cm.campaignid = t0.campaignid 
    LEFT JOIN wd_pa_programs2sites ps ON ps.p2sid = cm.p2sid 
    LEFT JOIN wd_pa_sites st ON st.siteid = ps.siteid LEFT JOIN countries ct ON ct.code = t0.countrycode 
    WHERE (t0.dateinserted BETWEEN '2013-1-1 00:00:00' AND '2013-12-31 23:59:59') AND t0.affiliateid = 'e7f63816' AND t0.rstatus IN (1, 2, 3) AND t0.transtype IN (1, 2, 4, 8, 16, 0) AND t0.hidden = 0 ORDER BY null ) 
    UNION 
    ( SELECT t0.transid as transid, 0 as transtype, t0.dateinserted as dateinserted, t0.countrycode as countrycode, t0.subnet_id as subnet_id, 0 as totalcost, t0.bannerid as bannerid, 2 as rstatus, t0.ip as ip, 0 as gender, 0 as orderid, 0 as payoutstatus, 0 as commission, 0 as currency_rate, 0 as parent_id, 0 as decline_reason, REPLACE( t1.name,".com","") as site, ct.name as countryname 
    FROM transactions_clicks_alcuda t0 
    INNER JOIN wd_pa_sites t1 ON t0.siteid = t1.siteid 
    LEFT JOIN countries ct ON ct.code = t0.countrycode 
    WHERE (t0.dateinserted BETWEEN '2013-1-1 00:00:00' AND '2013-12-31 23:59:59') AND t0.affiliateid = 'e7f63816' AND t0.user_data_id = 79838 ORDER BY null ) 
) ut LEFT JOIN users_subnets us ON us.id = ut.subnet_id 
ORDER BY dateinserted asc LIMIT 0, 20
;
