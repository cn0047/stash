Event Scheduler
-

Event Scheduler does not have any protection against multiple execution.

````sql
SHOW VARIABLES where Variable_name = 'event_scheduler';
SET GLOBAL event_scheduler = 1; -- must be root user

CREATE EVENT myevent
    ON SCHEDULE AT CURRENT_TIMESTAMP + INTERVAL 1 MINUTE
    DO
      CALL runScheduledEvent();

CREATE EVENT myevent
    ON SCHEDULE EVERY 60 SECOND
    DO
      CALL runScheduledEvent();

CREATE EVENT event1
  ON SCHEDULE EVERY 1 DAY STARTS '2019-07-11 01:00:00'
  DO SELECT NOW();

SHOW EVENTS;

DROP EVENT myevent;
````
