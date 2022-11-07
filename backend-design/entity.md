### Entity
- Users
    - id (int)
    - name (varchar)
    - occupation (varchar)
    - email (varchar)
    - password_hash (varchar)
    - avatar_file_name (varchar)
    - role (varchar)
    - token (varchar) --> menandai request dari frontend
    - created_at (datetime)
    - updated_at (datetime)

- Campaigns
    - id (int)
    - user_id (int)
    - name (varchar)
    - short_description (varchar)
    - description (text)
    - goal_amount (int)
    - current_amount (int)
    - backer_count (int)
    - slug (varchar)
    - created_at (datetime)
    - updated_at (datetime)

- Campaign Images
    - id (int)
    - campaign_id (int)
    - file_name (varchar)
    - is_primary boolean (small int)
    - created_at (datetime)
    - updated_at (datetime)

- Transactions
    - id (int)
    - campaign_id (int)
    - user_id (int)
    - amount (int)
    - created_at (datetime)
    - updated_at (datetime)

![erd image](/backend-design/erd.png "erd")