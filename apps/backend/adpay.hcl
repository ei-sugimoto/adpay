schema "public" {
  comment = "A schema comment"
}

schema "private" {}

table "users" {
  schema = schema.public
  column "id" {
    type = bigserial
  }
  column "name" {
    type = varchar(10)
  }
  column "password" {
    type = varchar(255)
  }

  primary_key {
    columns = [column.id]
  }

  index "idx_name" {
    columns = [column.name]
    unique = true
  }
}

table "projects" {
  schema = schema.public
  column "id" {
    type = bigserial
  }
  column "name" {
    type = varchar(10)
  }
  column "author_id" {
    type = bigint
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "fk_author"{
    columns = [column.author_id]
    ref_columns = [table.users.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }
}

table "bills" {
  schema = schema.public
  column "id" {
    type = bigserial
  }
  column "project_id" {
    type = bigint
  }
  column "amount" {
    type = int
  }

  column "src_user_id" {
    type = bigint
  }

  column "dst_user_id" {
    type = bigint
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "fk_project"{
    columns = [column.project_id]
    ref_columns = [table.projects.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  foreign_key "fk_src_user"{
    columns = [column.src_user_id]
    ref_columns = [table.users.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  foreign_key "fk_dst_user"{
    columns = [column.dst_user_id]
    ref_columns = [table.users.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }
}

table "project_permissions"{
  schema = schema.public
  column "id" {
    type = bigserial
  }
  column "project_id" {
    type = bigint
  }
  column "user_id" {
    type = bigint
  }

  primary_key {
    columns = [column.id]
  }

  foreign_key "fk_project"{
    columns = [column.project_id]
    ref_columns = [table.projects.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }

  foreign_key "fk_user"{
    columns = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update = NO_ACTION
    on_delete = CASCADE
  }
}
