create table product (
	id SERIAL primary key,
	product_name varchar(50) not null,
	price numeric(10,2) not null
);

select * from product;

insert into product (product_name, price) values ('Sushi', 100);
insert into product (product_name, price) values ('Batata Frita', 20);
insert into product (product_name, price) values ('Max Verstappen', 100000);
insert into product (product_name, price) values ('Carlos Sainz', 10000);
insert into product (product_name, price) values ('Charles Leclerc', 1000);
insert into product (product_name, price) values ('Checo', 0);

UPDATE product SET product_name='Sushi', price=20 WHERE id=8

