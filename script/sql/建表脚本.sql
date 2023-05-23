use test;

CREATE TABLE IF NOT EXISTS table_manage (
    id bigint(20) PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    c_table varchar(50) COMMENT '表名',
    c_table_name varchar(100) COMMENT '表名称',
    c_table_details varchar(2000) COMMENT '表说明',
    b_enable tinyint(1) COMMENT '是否可用 0：不可用 1：可用',
    c_create_by varchar(20) COMMENT '创建人',
    d_create_time datetime COMMENT '创建时间',
    c_update_by varchar(20) COMMENT '修改人',
    d_update_time datetime COMMENT '修改时间',
    c_delete_by varchar(20) COMMENT '删除人',
    d_delete_time datetime COMMENT '删除时间'
    ) COMMENT '表信息表';

CREATE TABLE IF NOT EXISTS field_manage (
    id bigint(20) PRIMARY KEY AUTO_INCREMENT COMMENT '主键',
    c_table_id bigint(20) NOT NULL COMMENT '表id，关联表：table_manage.id',
    c_field varchar(50) COMMENT '列名',
    c_field_name varchar(100) COMMENT '列名称',
    c_field_details varchar(2000) COMMENT '列说明',
    c_field_type varchar(30) COMMENT '列类型',
    i_field_length int COMMENT '字段长度',
    b_primary_key tinyint(1) COMMENT '是否为主键 0：非主键 1：主键',
    c_primary_rule varchar(20) COMMENT '主键生成规则 input：手动输入 auto_increment：自增长 snowflake_id：雪花ID',
    b_foreign_key tinyint(1) COMMENT '是否为外键 0：非外键 1：外键',
    c_association_table bigint(20) COMMENT '关联表',
    c_association_field bigint(20) COMMENT '关联字段',
    c_default_value varchar(2000) COMMENT '默认值',
    b_synchronous tinyint(1) COMMENT '是否已同步 0：未同步 1：已同步',
    b_enable tinyint(1) COMMENT '是否可用',
    c_create_by varchar(20) COMMENT '创建人',
    d_create_time datetime COMMENT '创建时间',
    c_update_by varchar(20) COMMENT '修改人',
    d_update_time datetime COMMENT '修改时间',
    c_delete_by varchar(20) COMMENT '删除人',
    d_delete_time datetime COMMENT '删除时间'
    ) COMMENT '字段信息表';