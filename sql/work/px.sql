-- arm db tracking
select
    us.*, ut.*
from userStat us
join urlTo ut on us.urlToHash = ut.hash
where us.userId = '69d43a269d5211e3ac4dd4bed9a9456d'
\G