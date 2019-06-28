Ruby on Rails
-
5.1.2

````sh
rails -v
rails new demo
rails server
````

## cli

````sh
rails generate controller Store index
````

````sh
rake db:seed
rake db:rollback
rake test
rake test:units
rake test:functionals
````

## db

````rb
order = Order.find(1)
order.pay_type = "Purchase order"
order.save

Order.where(name: 'dave').each do |order| puts order.amount
end
````

````sh
rails generate scaffold Product \
    title:string description:text image_url:string price:decimal

rails db:migrate RAILS_ENV=development
````

From book "Agile Web Development with Rails, 4th Edition, Rails 3.2" Finished chapters: 1, 2, 3, 5-10.
