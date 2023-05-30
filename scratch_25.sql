select count(*) from orders where customer_id = 1;
select name from customers;

explain analyze
select name, (
    select count(*)
    from orders
    where customer_id = customers.id
) orders_count from customers;

explain analyze
select name, count(*) orders_count
from customers
join orders o on customers.id = o.customer_id
group by customers.name;

select * from customers;
select * from orders;

explain analyze
select name from customers
where customers.id in (
    select customer_id
    from orders
    where total > 100
);

explain analyze
select distinct name
from customers
join orders o on customers.id = o.customer_id
where o.total > 100;

explain analyze
select name
from customers
join orders o on customers.id = o.customer_id
where o.total > 100
group by name;

select * from customers where id=1;

begin;
update customers set name = 'bud 2' where id = 1;
commit;
begin;
update customers set age = 100 where id = 1;
rollback;

describe customers;
alter table customers add index `cust_name_key` (name);
insert into customers (name) values
                                 ('andi'),
                                 ('budi'),
                                 ('budi'),
                                 ('bud1'),
                                 ('bud1'),
                                 ('bud0'),
                                 ('bud0'),
                                 ('kud0'),
                                 ('kud0'),
                                 ('kud0');

drop index cust_name_key on customers;

explain analyze
select * from customers
where name like 'ku%';