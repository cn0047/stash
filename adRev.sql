267698722   ->   270787129   ->   270789931   ->   389885287   ->   389895973   ->   267703520   ->   270892684
IWU         ->   SHGH        ->   UFI         ->   BN          ->   CL          ->   IWUCOUGAR   ->   BeD
            ->   Pay         ->   Pay         ->   Pay         ->               ->   Pay         ->   Pay
            ->   4.57        ->   4.57        ->   4.98        ->               ->   4.98        ->   4.98


select
request_uri
from profiles_http_referer where 
pid in (
267703520
);

select
  p.email
  ,phr.pid, phr.membership_since -- ,request_uri
  ,pre.external_profiles_registration_export_id eId
  ,ersI.site_name I__sourceSite -- partner_id
  ,epre.target_site_id I__target_site, epre.target_profile_id I__target_profile_id /*source_site_id, source_profile_id,*/
  ,ersE.site_name E__sourceSite -- partner_id
  ,epre.target_site_id E__target_site, epre.target_profile_id E__target_profile_id /*source_site_id, source_profile_id,*/
from profiles_http_referer phr
inner join profiles p on phr.pid = p.id
left JOIN external_profiles_registration_import epri ON epri.target_profile_id = phr.pid
left join external_registration_sites ersI on epri.source_site_id = ersI.id
left JOIN profiles_remarketing_external pre ON pre.base_profile_id = phr.pid
left JOIN external_profiles_registration_export epre ON epre.id = pre.external_profiles_registration_export_id
left join external_registration_sites ersE on epre.target_site_id = ersE.id
WHERE phr.visit_date between UNIX_TIMESTAMP(date(now()-interval 1 day)) and UNIX_TIMESTAMP(date(now()+interval 1 day))
and phr.request_uri LIKE '%0123456%'
;



select
phr.siteID
,phr.request_uri
,phr.pid
,epri.target_profile_id Import
,pre.base_profile_id Export
-- ,count(phr.pid)cnt
-- ,count(epri.target_profile_id)cntImport
-- ,count(pre.base_profile_id)cntExport
from profiles_http_referer phr
left JOIN external_profiles_registration_import epri ON epri.target_profile_id = phr.pid
left JOIN profiles_remarketing_external pre ON pre.base_profile_id = phr.pid
left JOIN external_profiles_registration_export epre ON epre.id = pre.external_profiles_registration_export_id
WHERE phr.visit_date between UNIX_TIMESTAMP(date(now()-interval 1 day)) and UNIX_TIMESTAMP(date(now()+interval 1 day))
and phr.request_uri LIKE '%0123456%'
-- group by phr.pid
;











===========================================================================================================================================================================
(IWU)267698722
(IWUCOUGAR)267703520

select * from profiles_http_referer where pid in (270787129 ) \G

select * from external_profiles_registration_import epri where epri.source_profile_id = 270787129;
select
    epri.source_site_id,
    epri.source_profile_id,
    epri.target_profile_id,
    epri.created_at,
    ers.site_name source_site_name,
    phr.request_uri
from external_profiles_registration_import epri
inner join external_registration_sites ers on epri.source_site_id = ers.id
left join profiles_http_referer phr on epri.target_profile_id = phr.pid
where epri.target_profile_id in (
267703520
);

select * from external_profiles_registration_export epre where epre.source_profile_id = 270787129;
select
    epre.source_site_id,
    epre.source_profile_id,
    epre.target_profile_id,
    epre.created_at,
    ers.site_name target_site_name
from external_profiles_registration_export epre
inner join external_registration_sites ers on epre.target_site_id = ers.id
where epre.source_profile_id in (
270787129
);


select
  p.email, o.siteID, o.user_id, o.id, epri.source_profile_id, epri.target_profile_id, ers.site_name source_site_name
from orders o
join profiles p on o.user_id = p.id
left join external_profiles_registration_import epri ON p.id = epri.target_profile_id
left join external_registration_sites ers on epri.source_site_id = ers.id
where o.id in (
  48032288, 48033016, 95320384, 42089050
);




===========================================================================================================================================================================


set sql_big_selects=1;
set @date1 = '2013-11-25';
set @date2 = '2013-11-28';
set @time1 = UNIX_TIMESTAMP(@date1);
set @time2 = UNIX_TIMESTAMP(@date2);
set @url = /*ON AM*/ 'http%3A%2F%2Fwww.iwantu.com%2Fextransfer.php%3Fdynamicpage%3D1click_full_cr%26source%3Dwww.iwantu.com%26transfer_to%3D102%26alc_cp%3D0000018308%26zoneid%3D1471';
set @url = /*ON GA*/ 'http%3A%2F%2Fwww.iwantu.com%2Fppc.php%3Fdynamicpage%3Dvideo_iwu%26ppc_cp%3D1234567890%26test_ad_rev%3D0123456';

-- regs
SELECT
  SUM(IF((lso1.order_id IS NULL OR ls1.id IS NULL OR o1.tdate < ls1.created_at) AND lro.order_id IS NULL AND o1.tdate <= @date2 AND
  pr.type = 1 AND traf.visit_date >= @time1 AND traf.visit_date <= @time2
  ,ROUND(o1.amount*IFNULL(oa1.revenue_rate,1),2),0)) AS revenue_regs_remarketing_type_1,
  SUM(IF((lso1.order_id IS NULL OR ls1.id IS NULL OR o1.tdate < ls1.created_at) AND lro.order_id IS NULL AND o1.tdate <= @date2 AND
  pr.type = 2 AND traf.visit_date >= @time1 AND traf.visit_date <= @time2
  ,ROUND(o1.amount*IFNULL(oa1.revenue_rate,1),2),0)) AS revenue_regs_remarketing_type_2,
  SUM(IF((lso1.order_id IS NULL OR ls1.id IS NULL OR o1.tdate < ls1.created_at) AND lro.order_id IS NULL AND o1.tdate <= @date2 AND
  pr.type = 3 AND traf.visit_date >= @time1 AND traf.visit_date <= @time2
  ,ROUND(o1.amount*IFNULL(oa1.revenue_rate,1),2),0)) AS revenue_regs_remarketing_type_3,
  SUM(IF((lso1.order_id IS NULL OR ls1.id IS NULL OR o1.tdate < ls1.created_at) AND lro.order_id IS NULL AND o1.tdate <= @date2 AND
  pr.type = 4 AND traf.visit_date >= @time1 AND traf.visit_date <= @time2
  ,ROUND(o1.amount*IFNULL(oa1.revenue_rate,1),2),0)) AS revenue_regs_remarketing_type_4,
  SUM(IF((lso1.order_id IS NULL OR ls1.id IS NULL OR o1.tdate < ls1.created_at) AND lro.order_id IS NULL AND o1.tdate <= @date2 AND
  pr.type = 5 AND traf.visit_date >= @time1 AND traf.visit_date <= @time2
  ,ROUND(o1.amount*IFNULL(oa1.revenue_rate,1),2),0)) AS revenue_regs_remarketing_type_5
FROM profiles_http_referer traf
LEFT JOIN profiles_remarketing pr ON pr.base_profile_id = traf.pid
LEFT JOIN external_profiles_registration_import epri ON epri.target_profile_id = traf.pid  
LEFT JOIN orders o1 ON (o1.user_id=pr.profile_id AND o1.apr_code IN ('AUTH_OK', 'CHARGED', 'STOPPED', 'PARTLY') AND o1.amount>0 AND o1.tdate <= @date2 AND pr.base_profile_id <> pr.profile_id )
LEFT JOIN orders_addons oa1 ON o1.id = oa1.order_id
LEFT JOIN orders_convertamount oc1 ON (UNIX_TIMESTAMP(DATE(o1.tdate)) = oc1.dt AND oc1.convert_currency = LEFT(o1.currency,3)) 
LEFT JOIN logins_by_source_orders lso1 ON (lso1.order_id = o1.id)
LEFT JOIN logins_by_source ls1 ON (ls1.id = lso1.logins_by_source_id AND ls1.track_type NOT IN (5,6))
LEFT JOIN logins_remarketing_orders lro ON (lro.order_id = o1.id) LEFT JOIN widget_users wu ON wu.profile_id = traf.pid 
WHERE visit_date >= @time1 AND visit_date <= @time2
and traf.request_uri = @url
;


-- logins
/*
explain
SELECT
SUM(IF(o.tdate <= @date2 AND
pr.type = 1 AND ls.created_at >= @time2 AND ls.created_at <= @time2
,ROUND(o.amount*IFNULL(oa.revenue_rate,1),2) ,0)) AS revenue_logins_remarketing_type_1,
SUM(IF(o.tdate <= @date2 AND
pr.type = 2 AND ls.created_at >= @time2 AND ls.created_at <= @time2
,ROUND(o.amount*IFNULL(oa.revenue_rate,1),2) ,0)) AS revenue_logins_remarketing_type_2,
SUM(IF(o.tdate <= @date2 AND
pr.type = 3 AND ls.created_at >= @time2 AND ls.created_at <= @time2
,ROUND(o.amount*IFNULL(oa.revenue_rate,1),2) ,0)) AS revenue_logins_remarketing_type_3,
SUM(IF(o.tdate <= @date2 AND
pr.type = 4 AND ls.created_at >= @time2 AND ls.created_at <= @time2
,ROUND(o.amount*IFNULL(oa.revenue_rate,1),2) ,0)) AS revenue_logins_remarketing_type_4,
SUM(IF(o.tdate <= @date2 AND
pr.type = 5 AND ls.created_at >= @time2 AND ls.created_at <= @time2
,ROUND(o.amount*IFNULL(oa.revenue_rate,1),2) ,0)) AS revenue_logins_remarketing_type_5
FROM logins_by_source ls
INNER JOIN logins_remarketing_orders lso ON (lso.logins_by_source_id = ls.id)
LEFT JOIN profiles_remarketing pr ON (pr.profile_id = lso.to_profile_id)
LEFT JOIN orders o ON (o.id=lso.order_id AND o.amount>0 AND o.tdate <= @date2 )
INNER JOIN profiles p ON (ls.profile_id=p.id AND if('' <> '', p.city = '', 1))
LEFT JOIN profiles_tester ptest ON (ptest.uid = p.id) 
LEFT JOIN orders_addons oa ON(o.id = oa.order_id)
-- LEFT JOIN currency_rates AS cr_gbp ON cr_gbp.date = DATE(o.tdate) AND cr_gbp.from = LEFT(o.currency, 3) AND cr_gbp.to = 'GBP'
-- LEFT JOIN currency_rates AS cr_usd ON cr_usd.date = DATE(o.tdate) AND cr_usd.from = LEFT(o.currency, 3) AND cr_usd.to = 'USD'
LEFT JOIN orders_convertamount oc ON (UNIX_TIMESTAMP(DATE(o.tdate)) = oc.dt AND oc.convert_currency = LEFT(o.currency,3))
INNER JOIN profiles_country pc FORCE INDEX(pid) ON (ls.profile_id=pc.pid AND pc.country_code3 IN ('gbr')) 
WHERE ls.created_at >= @date1 AND ls.created_at <= @date2
AND ls.track_type NOT IN (5,6)
AND ls.request_uri = @url
;
*/

-- regs ext
SELECT
traf.cache_request_uri AS tmp_hash,
p.id AS profile_id,
SUM(IF(expfr.logins_by_source_id IS NULL AND o.created_at <= @date2 AND
IF(o.remarketing_type > 0, o.remarketing_type,
CASE
WHEN exp.registration_type = 'banner' THEN 1
WHEN exp.registration_type = 'coregistration' THEN 2
WHEN exp.registration_type = 'posttransaction' THEN 3
WHEN exp.registration_type = 'crossmail' THEN 4
WHEN exp.registration_type = 'transfer' THEN 5
END) = 1 AND traf.visit_date >= @time1 AND traf.visit_date <= @time2
,o.amount,0)) AS revenue_regs_remarketing_type_1,
SUM(IF(expfr.logins_by_source_id IS NULL AND o.created_at <= @date2 AND
IF(o.remarketing_type > 0, o.remarketing_type,
CASE
WHEN exp.registration_type = 'banner' THEN 1
WHEN exp.registration_type = 'coregistration' THEN 2
WHEN exp.registration_type = 'posttransaction' THEN 3
WHEN exp.registration_type = 'crossmail' THEN 4
WHEN exp.registration_type = 'transfer' THEN 5
END) = 2 AND traf.visit_date >= @time1 AND traf.visit_date <= @time2
,o.amount,0)) AS revenue_regs_remarketing_type_2,
SUM(IF(expfr.logins_by_source_id IS NULL AND o.created_at <= @date2 AND
IF(o.remarketing_type > 0, o.remarketing_type,
CASE
WHEN exp.registration_type = 'banner' THEN 1
WHEN exp.registration_type = 'coregistration' THEN 2
WHEN exp.registration_type = 'posttransaction' THEN 3
WHEN exp.registration_type = 'crossmail' THEN 4
WHEN exp.registration_type = 'transfer' THEN 5
END) = 3 AND traf.visit_date >= @time1 AND traf.visit_date <= @time2
,o.amount,0)) AS revenue_regs_remarketing_type_3,
SUM(IF(expfr.logins_by_source_id IS NULL AND o.created_at <= @date2 AND
IF(o.remarketing_type > 0, o.remarketing_type,
CASE
WHEN exp.registration_type = 'banner' THEN 1
WHEN exp.registration_type = 'coregistration' THEN 2
WHEN exp.registration_type = 'posttransaction' THEN 3
WHEN exp.registration_type = 'crossmail' THEN 4
WHEN exp.registration_type = 'transfer' THEN 5
END) = 4 AND traf.visit_date >= @time1 AND traf.visit_date <= @time2
,o.amount,0)) AS revenue_regs_remarketing_type_4,
SUM(IF(expfr.logins_by_source_id IS NULL AND o.created_at <= @date2 AND
IF(o.remarketing_type > 0, o.remarketing_type,
CASE
WHEN exp.registration_type = 'banner' THEN 1
WHEN exp.registration_type = 'coregistration' THEN 2
WHEN exp.registration_type = 'posttransaction' THEN 3
WHEN exp.registration_type = 'crossmail' THEN 4
WHEN exp.registration_type = 'transfer' THEN 5
END) = 5 AND traf.visit_date >= @time1 AND traf.visit_date <= @time2
,o.amount,0)) AS revenue_regs_remarketing_type_5
from profiles_http_referer traf
INNER JOIN profiles_remarketing_external pre ON pre.base_profile_id = traf.pid
INNER JOIN external_profiles_registration_export exp ON exp.id = pre.external_profiles_registration_export_id
INNER JOIN external_registration_sites ers on exp.target_site_id = ers.id AND ers.partner_id in (3,1)
LEFT JOIN external_profiles_registrations_export_from_recovery expfr ON expfr.external_profiles_registration_export_id = exp.id
INNER JOIN external_remarketing_exported_profiles_orders o ON o.external_profiles_registration_export_id = exp.id AND o.created_at <= @date2 AND o.status IN ('CHARGED', 'STOPPED', 'PARTLY')
INNER JOIN profiles p ON (traf.pid = p.id AND if('' <> '', p.city = '', 1))
LEFT JOIN profiles_tester ptest ON (ptest.uid = p.id) 
LEFT JOIN currency_rates AS cr_gbp ON cr_gbp.date = DATE(o.created_at) AND cr_gbp.from = LEFT(o.currency, 3) AND cr_gbp.to = 'GBP'
LEFT JOIN currency_rates AS cr_usd ON cr_usd.date = DATE(o.created_at) AND cr_usd.from = LEFT(o.currency, 3) AND cr_usd.to = 'USD'
LEFT JOIN orders_convertamount oc ON (UNIX_TIMESTAMP(DATE(o.created_at)) = oc.dt AND oc.convert_currency = o.currency) 
WHERE visit_date >= @time1 AND visit_date <= @time2
AND (traf.request_uri LIKE @url)  
GROUP BY tmp_hash, visit_date
;