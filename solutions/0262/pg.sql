select
    t.request_at as "Day",
    round(1.0 * sum(
            case when t.status in ('cancelled_by_client', 'cancelled_by_driver') then
                1
            else
                0
            end) / count(*), 2) as "Cancellation Rate"
from
    trips t
    join users uc on t.client_id = uc.users_id
        and uc.banned = 'No'
    join users ud on t.driver_id = ud.users_id
        and ud.banned = 'No'
where
    t.request_at between '2013-10-01' and '2013-10-03'
group by
    "Day";
