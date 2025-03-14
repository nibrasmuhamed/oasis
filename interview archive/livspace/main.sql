-- table : city
-- two columns; city_id, name

-- table : customer
-- three colums: customer_id, name, city_id

-- table : order
-- three columns: id, value, customer_id

-- in every city as cityname, number of hni
-- hni is individual spend more than 50k in entire time.
Select c.name as cityname, COUNT(h.hnicustomerid) as hni
from (
    select customer_id as hnicustomerid 
    from order 
    group by customer_id 
    having sum(value) > 50000;
) as h
JOIN customer AS cust on h.hnicustomerid = cust.customer_id
JOIN city as c on cust.city_id = c.city_id
Group by c.city_id;


-- city_id : 1
-- name : wayanad 
-- value: 100

-- city_id: 2
-- name : blr
-- value: 200

-- city_id: 2
-- name : blr
-- value: 300

-- huge file of 1gb, it has only one columns which has numbers. 
-- sort this file. we have only 100mb available