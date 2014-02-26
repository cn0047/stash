------------------------------------------------------------------------------------------------------------------------
-- TRACKING
------------------------------------------------------------------------------------------------------------------------

-- История переданых в афф систему транзакций за 7 дней недели
set @v_like  = '%trackingMethod=5%';
explain select qs_main, qs_additional
    from tracking_codes_posted_history_1
    where qs_main like @v_like or qs_additional like @v_like
    limit 1
union all
    select qs_main, qs_additional
    from tracking_codes_posted_history_2
    where qs_main like @v_like or qs_additional like @v_like
    limit 1
union all
    select qs_main, qs_additional
    from tracking_codes_posted_history_3
    where qs_main like @v_like or qs_additional like @v_like
    limit 1
union all
    select qs_main, qs_additional
    from tracking_codes_posted_history_4
    where qs_main like @v_like or qs_additional like @v_like
    limit 1
union all
    select qs_main, qs_additional
    from tracking_codes_posted_history_5
    where qs_main like @v_like or qs_additional like @v_like
    limit 1
union all
    select qs_main, qs_additional
    from tracking_codes_posted_history_6
    where qs_main like @v_like or qs_additional like @v_like
    limit 1
union all
    select qs_main, qs_additional
    from tracking_codes_posted_history_7
    where qs_main like @v_like or qs_additional like @v_like
    limit 1
;

-- Add tracking codes SETTINGS
insert into tracking_codes_settings set class='main', name='ProductID', site_id=51, sub_site_id=0, use_for='product',             value='dms_comP';
insert into tracking_codes_settings set class='main', name='ProductID', site_id=51, sub_site_id=0, use_for='male_confirmation',   value='dms_comM';
insert into tracking_codes_settings set class='main', name='ProductID', site_id=51, sub_site_id=0, use_for='female_confirmation', value='dms_comF';
insert into tracking_codes_settings set class='main', name='ProductID', site_id=51, sub_site_id=0, use_for='couple_confirmation', value='dms_comC';

-- Add TRACKING CODES
select * from tracking_codes where site_id in (36) and title like 'AFF%';
insert into tracking_codes select 
    null,group_id_fk,title,code,code_for_payer,code_for_orders,code_for_genders,code_for_age1,code_for_age2,code_for_packages,class,source,network_id_fk,date_created,display_position,display_mode,display_use,amount_minimal,daily_cap,confirmed_registration_limit,confirmed_days_limit,tracking_param,tracking_value,
    51, -- site id
    country_code,region_code,is_debug_allowed,is_unique_registration_required,do_not_track,updated,do_not_track_female,funnel_tracking,target_roi,track_3d_pay,save_to_aff_sys
from tracking_codes where code_id in (19567);

-- История переданых кликов
select * from tracking_codes_posted_click_history_3 where qs_main like '%a_aid=46b27c3f%';

-- Клик в базе
select * from tracking_codes_posted_click_0 where qs_main like '%a_aid=46b27c3f%' limit 1;
select * from tracking_codes_posted_click_1 where qs_main like '%a_aid=46b27c3f%' limit 1;
select * from tracking_codes_posted_click_2 where qs_main like '%a_aid=46b27c3f%' limit 1;
select * from tracking_codes_posted_click_3 where qs_main like '%a_aid=46b27c3f%' limit 1;
select * from tracking_codes_posted_click_4 where qs_main like '%a_aid=46b27c3f%' limit 1;
select * from tracking_codes_posted_click_5 where qs_main like '%a_aid=46b27c3f%' limit 1;

-- tracking codes posted
--  CLICS
select
    (select count(*) from tracking_codes_posted_click_0) as t0,
    (select count(*) from tracking_codes_posted_click_1) as t1,
    (select count(*) from tracking_codes_posted_click_2) as t2,
    (select count(*) from tracking_codes_posted_click_3) as t3,
    (select count(*) from tracking_codes_posted_click_4) as t4,
    (select count(*) from tracking_codes_posted_click_5) as t5
;
--  REG + CONF + PAY
select
    (select count(*) from tracking_codes_posted_0) as t0,
    (select count(*) from tracking_codes_posted_1) as t1,
    (select count(*) from tracking_codes_posted_2) as t2,
    (select count(*) from tracking_codes_posted_3) as t3,
    (select count(*) from tracking_codes_posted_4) as t4,
    (select count(*) from tracking_codes_posted_5) as t5
;
--  CONVENIENT REG + CONF + PAY
select * from (
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_0 union all
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_1 union all
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_2 union all
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_3 union all
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_4 union all
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_5
)t where qs_additional like '%a_aid%' and profile_id_fk = 269403328
;
-- history
select
    (select count(*) from tracking_codes_posted_history_1) as t1,
    (select count(*) from tracking_codes_posted_history_2) as t2,
    (select count(*) from tracking_codes_posted_history_3) as t3,
    (select count(*) from tracking_codes_posted_history_4) as t4,
    (select count(*) from tracking_codes_posted_history_5) as t5,
    (select count(*) from tracking_codes_posted_history_6) as t6,
    (select count(*) from tracking_codes_posted_history_7) as t7
;
-- ID in history
select *from (
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_history_1 union all
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_history_2 union all
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_history_3 union all
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_history_4 union all
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_history_5 union all
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_history_6 union all
  select site_id, code_id_fk, profile_id_fk, order_id_fk, qs_main, qs_additional from tracking_codes_posted_history_7
)t where qs_additional like '%a_aid%' and profile_id_fk = 270642706
;
-- MOVE from history to queue
insert into tracking_codes_posted_5 select
    null, code_id_fk, profile_id_fk, order_id_fk, order_type, url, qs_main, qs_additional, posted, sent, site_id
from tracking_codes_posted_history_5;

set @pid=216869918;
          select FROM_UNIXTIME(posted),code_id_fk,profile_id_fk pid,order_id_fk,order_type t,qs_main, qs_additional,site_id from tracking_codes_posted_0 where profile_id_fk=@pid
union all select FROM_UNIXTIME(posted),code_id_fk,profile_id_fk pid,order_id_fk,order_type t,qs_main, qs_additional,site_id from tracking_codes_posted_1 where profile_id_fk=@pid
union all select FROM_UNIXTIME(posted),code_id_fk,profile_id_fk pid,order_id_fk,order_type t,qs_main, qs_additional,site_id from tracking_codes_posted_2 where profile_id_fk=@pid
union all select FROM_UNIXTIME(posted),code_id_fk,profile_id_fk pid,order_id_fk,order_type t,qs_main, qs_additional,site_id from tracking_codes_posted_3 where profile_id_fk=@pid
union all select FROM_UNIXTIME(posted),code_id_fk,profile_id_fk pid,order_id_fk,order_type t,qs_main, qs_additional,site_id from tracking_codes_posted_4 where profile_id_fk=@pid
union all select FROM_UNIXTIME(posted),code_id_fk,profile_id_fk pid,order_id_fk,order_type t,qs_main, qs_additional,site_id from tracking_codes_posted_5 where profile_id_fk=@pid;

truncate tracking_codes_posted_click_0;
truncate tracking_codes_posted_click_1;
truncate tracking_codes_posted_click_2;
truncate tracking_codes_posted_click_3;
truncate tracking_codes_posted_click_4;
truncate tracking_codes_posted_click_5;

truncate tracking_codes_posted_0;
truncate tracking_codes_posted_1;
truncate tracking_codes_posted_2;
truncate tracking_codes_posted_3;
truncate tracking_codes_posted_4;
truncate tracking_codes_posted_5;

truncate tracking_codes_posted_history_1;
truncate tracking_codes_posted_history_2;
truncate tracking_codes_posted_history_3;
truncate tracking_codes_posted_history_4;
truncate tracking_codes_posted_history_5;
truncate tracking_codes_posted_history_6;
truncate tracking_codes_posted_history_7;

-- PAYments tracking.
select
    o.siteID,o.SubSiteID,o.id,o.user_id,o.tdate,o.amount,o.currency,o.product
    ,c.real_rate
    ,round(o.amount/c.real_rate, 2) trackAmount
    ,ifnull(t.decline_reason, '')status,ifnull(t.code_id,'')
    ,if(l.logins_by_source_id is not null, 'yes', '')login
from orders o 
left join tracking_codes_processed_additional t on o.user_id = t.profile_id and o.id = t.order_id
left join logins_by_source_orders l on o.id = l.order_id
left join currency_rates c on date(o.tdate) = c.date and c.to = 'GBP' and c.from = left(o.currency, 3)
where o.id in (
    48266964
);

-- Проверяет зашел ли трек в базу
set @v_pid  = '217701504';
select * from tracking_codes_posted_5 where profile_id_fk = @v_pid;

set @v_aid  = '%e7f63816%';
select * from tracking_codes_posted_9 where qs_additional like @v_aid or qs_main like @v_aid;

select 
    tcph. id, profile_id_fk, order_id_fk, qs_additional, posted, sent,
    tcpl. ids, hour, processed
from 
    tracking_codes_posted_5 tcph 
    left join tracking_codes_posted_log tcpl on tcph.id = tcpl.ids
where 
    tcph.profile_id_fk = @v_pid;

-- Проверяет был ли трек передан в AFFILIATES SYSTEM 
-- по id профайла
set @v_pid  = '217701504';
select 
    id, profile_id_fk, order_id_fk, qs_main, qs_additional, posted, sent 
from tracking_codes_posted_history_5 where profile_id_fk = @v_pid;
-- по a_aid аффилейта
set @v_like  = '%e7f63816%';
select 
    id, profile_id_fk, order_id_fk, qs_main, qs_additional, posted, sent 
from tracking_codes_posted_history_1 where qs_additional like @v_like;


-- ORDER Смотрим ордер по его id и URL по которому пришел юзер
select o. id, user_id, phr. request_uri, http_referer
from orders o 
left join profiles_http_referer phr on o.user_id = phr.pid
where o.id in (41966607, 41966615);


-- Лог отправки треков в аффилейт систему
set @v_uts  = 1345683600;
select action, hour, sum(ids) from tracking_codes_posted_log where processed > @v_uts group by action, hour order by hour, action;
select action, sum(ids), group_concat(hour) from tracking_codes_posted_log where processed > @v_uts group by action order by hour, action;
select action, hour, sum(ids) from tracking_codes_posted_log where processed > @v_uts and hour <> 23 group by action, hour order by hour, action;
select action, sum(ids), group_concat(hour) from tracking_codes_posted_log where processed > @v_uts and hour <> 23 group by action order by hour, action;


-- Кол-во треков с и без order_id_fk
select sum(if(order_id_fk > 0,1,0)) with_order, sum(if(order_id_fk <= 0,1,0)) without_order from tracking_codes_posted_7;
select * from tracking_codes_posted_7 where order_id_fk <= 0 limit 3;









------------------------------------------------------------------------------------------------------------------------
-- TRANSACTIONS HISTORY
------------------------------------------------------------------------------------------------------------------------
/*
    base multi_c table 
    tracking_codes_processed
    contains declined trackCodes
*/

-- sale debug Кол-во полученных транзакции по часам
select HOUR(FROM_UNIXTIME(sent)) as hour, count(sent)  from tracking_codes_posted_history_4 group by hour;
-- sale debug Кол-во полученных кликов по часам
select HOUR(FROM_UNIXTIME(sent)) as hour, count(sent)  from tracking_codes_posted_click_history_3 group by hour;

-- выгребем ордеры полученных платежей, кроме тех что есть в аффилейт системе
SELECT GROUP_CONCAT(order_id_fk SEPARATOR "','") as ordrs
FROM tracking_codes_posted_history_5 
WHERE 
    qs_main LIKE '%_comP%' 
    AND  order_id_fk NOT IN ('42029475','42029483','42029573','42029597','42029670','42029730','42029776','42029840','42029889','42029913','42029939','42030052','42030031','42030200','42030240','42030258','42030487','42030531','42030574','42030705','42030808','42030815','42030952','42031675','42031725','42031735','42031878','42032004','42032176','42032465','41981910','41982549','41983247','42020634','41991840','41996041','42000212','42026450','42019292','42009101','42008394','42024352','42032898','42032899','42032901','42032994','42033014','42033046','42033137','42033291','42033348','42033372','42033388','42033426','42033554','42033928','42034165','42034305','42034317','42034324','42034437','42034750','42034832','42034834','42034837','42034878','42035097','42035181','42035213','42024330','42029326','42035215','42035221','42035312','42035359','42035410','42035418','42035473','42035630','42035652','42035730','42035747','42035757')
;
-- LIMIT 20;

-- кол-во полученных транзакции в разрезе по аффилейтам
select RIGHT(LEFT(qs_additional, LOCATE('a_aid=', qs_additional)+13), 8) as a_aid, count(qs_additional) 
from tracking_codes_posted_history_4
group by a_aid order by a_aid ;

select count(*) from tracking_codes_posted_history_2 where qs_additional like '%d581b204%' and FROM_UNIXTIME(sent) > '2012-09-11 01:00:00';



------------------------------------------------------------------------------------------------------------------------

-- tracking_codes_processed -- contains do all (and not tracked too) transactions
-- profiles_http_referer
-- logins_by_source

-- Show count of users by reg-type and aff
select date(FROM_UNIXTIME(phr.confirmation_date)) d, p.reg_type rt, count(p.id) cnt
from profiles_http_referer phr
join profiles p on phr.pid=p.id
where phr.siteID=30 and phr.confirmation_date>UNIX_TIMESTAMP('2013-01-05 01:01:01') and phr.request_uri like '%17691891%'
and p.siteID=30 and p.reg_type in (24,25)
group by d,rt;



-- SHOWS NOT TRACKED AFFILIATES PAYMENTS
-- explain
set @date1 = date(now()-interval 12 day);
set @date2 = date(now()-interval 1 day);
set @site = 30;
set SQL_BIG_SELECTS=1;
select * from ( -- select count(*) from (
    select
        o.id oid, o.user_id, o.amount -- tdate, user_id, o.siteID, amount, `type`, type_payment, apr_code, product
        -- p.siteID,p.SubSiteID,p.screenname,p.smspin,p.id,p.xid,p.type_id,p.photo,p.lname,p.password,p.email,p.phone_number,p.mobile_number,p.city,p.city_id,p.state,p.zip,p.country,p.lastvisit,p.membership_since,p.membership_type,p.subscriber,p.membership_expire,p.user_key,p.user_code,p.confirmation_date,p.is_vip,p.confirmed,p.searchable,p.blocked,p.approved,p.reg_type,p.fname,p.emailstatus,p.reference,p.getyounoticed_,p.accomodate,p.travel,p.video,p.video_expire,p.video_subscriber,p.video_reference,p.is_failed,p.updated,p.with_sms,p.countryfromip,p.osx,p.osy
        -- , p.country -- membership_type-- ,p.reg_type -- , email, is_vip, blocked, approved, is_failed
        , tcp.order_id Toid, tcpa.order_id Toida
        ,datediff(now(), phr.membership_since) D, phr.source/*, http_referer*/, right(left(phr.request_uri, locate('a_aid', phr.request_uri)+15), 8) a_aid
        -- , case when oa.payment_source is null then 'web' when oa.payment_source=1 then 'web'  when oa.payment_source=2 then 'mob' else oa.payment_source end place
        -- select count(*)
        -- , if(phr.request_uri like '%a_bid%', '+', '-') abid -- , right(left(phr.request_uri, locate('a_bid', phr.request_uri)+15), 8) a_bid
        -- ,phr.request_uri
        -- ,mdl.device_os,device_os_version,model_name,brand_name
        ,ls.profile_id LOGIN
    -- select count(*)
    from orders o
    -- left join profiles p on o.user_id=p.id and p.siteID=@site
    -- join orders_addons oa on o.id=oa.order_id
    join profiles_http_referer phr on o.user_id=phr.pid and phr.request_uri like '%a_aid%'
    left join tracking_codes_processed tcp             on o.user_id=tcp.profile_id  and o.id=tcp.order_id  and tcp.scheme=3
    left join tracking_codes_processed_additional tcpa on o.user_id=tcpa.profile_id and o.id=tcpa.order_id and tcpa.scheme=3
    left join logins_by_source ls on o.user_id=ls.profile_id
    -- left join profiles_mobile_device pmd on o.user_id=pmd.profile_id
    -- left join mobile_devices_list mdl on pmd.device_id=mdl.id
    where o.siteID=@site and o.apr_code in ('CHARGED') and o.tdate between @date1 and @date2 and o.`type` in ('O', 'A')  and o.product not like '%3 Days Trial%'
    -- and phr.membership_since>'2013-05'
    -- and p.membership_type=0 and o.amount>10
    -- order by mobile_number -- place, order_id, order_ida
    -- limit 25;
-- )t where order_id is not null and order_ida is not null
)t where Toid is null and Toida is null and LOGIN is null and amount>10 and D<=30
limit 111
;


-- CONF
select
    p.id,p.confirmation_date,
    tcp.profile_id, tcpa.profile_id
    -- select count(distinct p.id)
from profiles p
join profiles_http_referer phr on p.id=phr.pid and phr.request_uri like '%a_aid%'
left join tracking_codes_processed tcp             on p.id=tcp.profile_id  and tcp.scheme in (1,2)
left join tracking_codes_processed_additional tcpa on p.id=tcpa.profile_id and tcpa.scheme in (1,2)
where p.siteid in (@site) and p.confirmation_date between @date1 and @date2
-- and phr.membership_since between @date1 and @date2
-- and p.type_id = 1
and tcp.profile_id is null and tcpa.profile_id is null
;




-- Shows count of cofirmed users more than once
set @date1 = date(now()-interval 2 day);
set @date2 = date(now()-interval 1 day);
select mt1.site,mt1.all_count_of_confirms,mt2.cofirmed_more_than_once
from (
    select site,count(site) all_count_of_confirms from (
        select case when t1.request_uri regexp 'amissexy'  then 'amissexy'
                  when t1.request_uri regexp 'milfberry' then 'milfberry'
                  when t1.request_uri regexp 'upforit'   then 'upforit'
                  else 'else'
            end site
        from profiles_http_referer t1
        join tracking_codes_processed t2 on t1.pid=t2.profile_id and t2.scheme=2
        where t1.siteID in (30,12) and t1.confirmation_date between unix_timestamp(@date1) and unix_timestamp(@date2) and t1.request_uri regexp '(amissexy|milfberry|upforit).*a_aid'
        and t2.date between date(@date1) and date(@date2)
    ) t_2 group by site
) mt1
left join (
select site,count(site) cofirmed_more_than_once from (
    select t3.profile_id,count(t3.profile_id)cnt,t3.date,t1.request_uri
        ,case when t1.request_uri regexp 'amissexy'  then 'amissexy'
              when t1.request_uri regexp 'milfberry' then 'milfberry'
              when t1.request_uri regexp 'upforit'   then 'upforit'
              else 'else'
        end site
        from profiles_http_referer t1
        join tracking_codes_processed t2 on t1.pid=t2.profile_id and t2.scheme=2
        join tracking_codes_processed_additional t3 on t1.pid=t3.profile_id and t3.scheme=2
        where t1.siteID in (30,12) and t1.confirmation_date between unix_timestamp(@date1) and unix_timestamp(@date2) and t1.request_uri regexp '(amissexy|milfberry|upforit).*a_aid'
        and t2.date between date(@date1) and date(@date2)
        and t3.date between date(@date1) and date(@date2)
        group by t3.profile_id
        having cnt>1
        order by cnt
    ) t_1 group by site
) mt2
on mt1.site=mt2.site;








------------------------------------------------------------------------------------------------------------------------
-- EXTERNAL REMARKETING
------------------------------------------------------------------------------------------------------------------------
set sql_big_selects=1;
set @date1 = '2013-12-01';
set @date2 = '2013-12-10';

-- Show GOOD external orders per day.
select
  date(created_at)d,count(date(created_at))cnt,sum(amount)amount
from external_remarketing_exported_profiles_orders 
where created_at between @date1 and @date2 and status in ('CHARGED', 'STOPPED', 'PARTLY')
group by d
;
-- Show REGS & GOOD external orders per country.
SELECT
  ers.partner_id, pc.country_code3, count(*)regs, count(o.amount)payCount, sum(o.amount)amount
FROM profiles_http_referer traf
INNER JOIN profiles p ON (traf.pid = p.id AND if('' <> '', p.city = '', 1))
INNER JOIN profiles_country pc ON traf.pid = pc.pid
INNER JOIN profiles_remarketing_external pre ON pre.base_profile_id = traf.pid
INNER JOIN external_profiles_registration_export exp ON exp.id = pre.external_profiles_registration_export_id
INNER JOIN external_registration_sites ers on exp.target_site_id=ers.id
LEFT JOIN external_profiles_registrations_export_from_recovery expfr ON expfr.external_profiles_registration_export_id = exp.id
LEFT JOIN external_remarketing_exported_profiles_orders o ON o.external_profiles_registration_export_id = exp.id AND o.created_at <= @date2 AND o.status IN ('CHARGED', 'STOPPED', 'PARTLY')
WHERE visit_date >= UNIX_TIMESTAMP(@date1) AND visit_date <= UNIX_TIMESTAMP(@date2) /*AND traf.siteID IN (30)*/
GROUP BY ers.partner_id, pc.country_code3
ORDER BY ers.partner_id, regs DESC
LIMIT 11
;
-- External RAW REVENUE per day and PARTNER ID.
SELECT
  date(o.created_at)d,
  count(date(o.created_at))totalCnt,
  sum(o.amount)totalAmount,
  sum(if (ers.partner_id = 3, 1, 0)) casCnt,
  sum(if (ers.partner_id = 3, o.amount, 0)) casAmount,
  sum(if (ers.partner_id = 1, 1, 0)) ourCnt,
  sum(if (ers.partner_id = 1, o.amount, 0)) ourAmount
FROM profiles_http_referer traf
INNER JOIN profiles_remarketing_external pre ON traf.pid = pre.base_profile_id
INNER JOIN external_profiles_registration_export exp ON pre.external_profiles_registration_export_id = exp.id
INNER JOIN external_registration_sites ers on exp.target_site_id = ers.id
LEFT JOIN external_remarketing_exported_profiles_orders o ON exp.id = o.external_profiles_registration_export_id AND o.created_at <= @date2 AND o.status IN ('CHARGED', 'STOPPED', 'PARTLY')
WHERE o.created_at between @date1 and @date2
group by d
;
-- [+] Shows REGS & RAW REVENUE by sites and locations for IMPORT by PARTNER ID.
select
    ers.site_name,
    pc.country_code3,
    count(ers.site_name)regs, sum(o.amount)RAWRevenue
from external_profiles_registration_import epri
join external_registration_sites ers on epri.source_site_id = ers.id
join profiles_country pc on epri.target_profile_id = pc.pid
left join orders o on o.user_id = epri.target_profile_id
where epri.created_at between @date1 and @date2
and ers.partner_id in (1)
group by ers.site_name, pc.country_code3
order by ers.site_name, regs desc
;
-- [+] Shows RAW REVENUE by sites and locations and source sites adn PARTNER ID..
SELECT
    ers.partner_id, ers.site_name,
    traf.siteID,
    pc.country_code3,
    min(date(traf.membership_since))regDate,
    min(date(o.created_at))payDate, sum(o.amount)RAWRevenue
FROM profiles_http_referer traf
INNER JOIN profiles p ON (traf.pid = p.id AND if('' <> '', p.city = '', 1))
INNER JOIN profiles_country pc ON traf.pid = pc.pid
INNER JOIN profiles_remarketing_external pre ON pre.base_profile_id = traf.pid
INNER JOIN external_profiles_registration_export exp ON exp.id = pre.external_profiles_registration_export_id
INNER JOIN external_remarketing_exported_profiles_orders o ON o.external_profiles_registration_export_id = exp.id AND o.created_at <= @date2 AND o.status IN ('CHARGED', 'STOPPED', 'PARTLY')
INNER JOIN external_registration_sites ers on exp.target_site_id=ers.id
LEFT JOIN external_profiles_registrations_export_from_recovery expfr ON expfr.external_profiles_registration_export_id = exp.id
WHERE visit_date >= UNIX_TIMESTAMP(@date1) AND visit_date <= UNIX_TIMESTAMP(@date2)
and ers.partner_id in (1)
GROUP BY ers.partner_id, ers.site_name, traf.siteID, pc.country_code3
ORDER BY traf.siteID, pc.country_code3, ers.site_name DESC
;
-- All about particular external PAY.
select
    ers.partner_id, ers.site_name,
    traf.siteID,
    pc.country_code3,
    traf.membership_since,
    o.created_at, o.amount
FROM profiles_http_referer traf
INNER JOIN profiles p ON (traf.pid = p.id AND if('' <> '', p.city = '', 1))
INNER JOIN profiles_country pc ON traf.pid = pc.pid
INNER JOIN profiles_remarketing_external pre ON pre.base_profile_id = traf.pid
INNER JOIN external_profiles_registration_export exp ON exp.id = pre.external_profiles_registration_export_id
LEFT JOIN external_profiles_registrations_export_from_recovery expfr ON expfr.external_profiles_registration_export_id = exp.id
LEFT JOIN external_remarketing_exported_profiles_orders o ON o.external_profiles_registration_export_id = exp.id
INNER JOIN external_registration_sites ers on exp.target_site_id=ers.id
where o.external_order_id in (
48048921
);

php cron/external_remarketing_stat_collect.php -a payments_info -f '2013-11-21' -t '2013-11-22'


------------------------------------------------------------------------------------------------------------------------
-- Shows when host used on registration do not equal to host in EDGE CASE.
------------------------------------------------------------------------------------------------------------------------
set sql_big_selects=1;
set @date1 = '2014-01-01';
set @date2 = '2014-02-20';
-- explain
select
    count(*)
    -- phr.siteID, ss.value, phr.pid, phr.request_uri
from profiles_http_referer phr
join site_settings ss on phr.siteID = ss.siteID and ss.name = 'MAIL_SITE_NAME' and ss.value <> 'White Label Dating'
where visit_date >= UNIX_TIMESTAMP(@date1) AND visit_date <= UNIX_TIMESTAMP(@date2)
-- and phr.siteID=57
and phr.request_uri like 'EDGE+CASE%7Chttp%'
and phr.request_uri not regexp ss.value
;