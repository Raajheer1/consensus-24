INSERT INTO accounts (email, first_name, last_name, company_name, profile_pic, public_key, created_at, updated_at)
VALUES ('the@raajpatel.dev', 'Raaj', 'Patel', 'Raaj Inc.', 'https://avatars.githubusercontent.com/u/47269252?v=4', 'ABCDfdsajkdjfskafklas', NOW(), NOW());

INSERT INTO accounts (email, first_name, last_name, company_name, profile_pic, public_key, created_at, updated_at)
VALUES ('loan@melby.social', 'Andrew', 'Melbourne', 'Andrew Loans', 'https://avatars.githubusercontent.com/u/47269252?v=4', 'FDSJFSDKLDFSJFDSLFKLJS', NOW(), NOW());

INSERT INTO loans (borrower_id, goal_amount, amount_raised, number_of_payments, payment_schedule, interest_rate, title, description, image_url, loan_token_asset_code, funded_at, created_at, updated_at)
VALUES (1, 5000.00, 0.00, 12, 'monthly', 3.75, 'New Mower for Growing Business', 'Looking to expand my lawn care business with a new commercial-grade zero-turn mower to handle larger properties and increase efficiency. This loan will help me purchase the mower, leading to more jobs and faster turnaround times for my customers.', 'https://www.shutterstock.com/image-photo/lawn-mover-on-green-grass-260nw-1730029327.jpg', 'HdCjjAWYLFy2BF4q1POIy5h5esspW2gRQPYe3wFE5MnsiXDAAh0iprs0', NULL, NOW(), NOW());

INSERT INTO loans (borrower_id, goal_amount, amount_raised, number_of_payments, payment_schedule, interest_rate, title, description, image_url, loan_token_asset_code, funded_at, created_at, updated_at)
VALUES (1, 2000.00, 0.00, 6, 'monthly', 5.99, 'Upgrade Edging Equipment', 'My current lawn equipment is showing its age. I need a loan to replace my aging trimmer, edger, and blower with newer, more reliable models. This investment will improve the quality of my service, reduce downtime, and keep my customers happy.', 'https://www.shutterstock.com/image-photo/process-cutting-lawn-cordless-grass-260nw-2166141793.jpg', 'zxOkHZIwcFrmFQkw3HD9EqiGMP6gFWGtvMyEug7k84DrKSSyLm8ZZZhI', NULL, NOW(), NOW());

INSERT INTO loans (borrower_id, goal_amount, amount_raised, number_of_payments, payment_schedule, interest_rate, title, description, image_url, loan_token_asset_code, funded_at, created_at, updated_at)
VALUES (1, 200.00, 0.00, 2, 'weekly', 1.99, 'Gas Money', 'I temporarily need gas money for my lawn mower business to get me through this week.', 'https://www.shutterstock.com/image-photo/process-cutting-lawn-cordless-grass-260nw-2166141793.jpg', 'zxOkHZIwcFrmFQkw3HD9EqiGMPFDDWGtvMyEug7k84DrKSSyLm8ZZZhI', NULL, NOW(), NOW());

INSERT INTO loan_lenders (loan_id, lender_id, loan_amount, created_at, updated_at)
VALUES (1, 2, 1000.00, NOW(), NOW());

INSERT INTO loan_lenders (loan_id, lender_id, loan_amount, created_at, updated_at)
VALUES (2, 2, 250.00, NOW(), NOW());

INSERT INTO feeds (loan_id, feed_type, title, description, created_at, updated_at, created_by)
VALUES (3, 'loan', 'Loan Created', 'I created this loan', NOW(), NOW(), 1);

INSERT INTO feeds (loan_id, feed_type, title, description, created_at, updated_at, created_by)
VALUES (3, 'loan', 'Update', 'Im in desperate need of gas!', NOW(), NOW(), 1);