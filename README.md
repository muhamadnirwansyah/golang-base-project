# golang base project

1. database -> postgresql
2. query -> goqu (golang query)
3. json web tokens

# structure project

<img width="255" alt="image" src="https://github.com/user-attachments/assets/a644a0b8-617e-43e8-b1a3-591e3c251e37" />


# how to use

1. inject database

  ```
    CREATE TABLE public.accounts (
	  id bigserial NOT NULL,
	  full_name varchar(200) NOT NULL,
	  email varchar(200) NOT NULL,
	  phone_number varchar(20) NOT NULL,
	  "password" varchar(255) NOT NULL,
	  CONSTRAINT accounts_email_key UNIQUE (email),
	  CONSTRAINT accounts_phone_number_key UNIQUE (phone_number),
	  CONSTRAINT accounts_pkey PRIMARY KEY (id));
  ```

2. register data with postman

<img width="857" alt="image" src="https://github.com/user-attachments/assets/2e2779cf-dbe1-41fb-a0a2-bea9a6ab123a" />

3. login

<img width="848" alt="image" src="https://github.com/user-attachments/assets/a6ae9d23-2e8e-4a8d-8945-efbf52625ee8" />

4. validation token

<img width="851" alt="image" src="https://github.com/user-attachments/assets/b24f186e-79bb-4ed0-b9ea-3a867e78a135" />


   
