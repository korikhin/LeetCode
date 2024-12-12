select
    t.request_at as "Day",
    round(1.0 * sum(
            case when t.status in ('cancelled_by_client', 'cancelled_by_driver') then
                1
            else
                0
            end) / count(*), 2) as "Cancellation Rate"
from
    public.trips as t
where
    exists (
        select
            null
        from
            public.users as u
        where
            u.users_id in (t.client_id, t.driver_id)
            and u.banned = 'No'
        having
            count(*) = 2)
    and t.request_at between '2013-10-01' and '2013-10-03'
group by
    "Day";
