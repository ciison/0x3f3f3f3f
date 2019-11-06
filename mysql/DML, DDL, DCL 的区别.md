## DML, DDL, DCL 的区别

* DML 

  * DML(data manipulation language) 数据操作语言, `SELECT`, `UPDATE`, `INSERT`, `DELETE` 。主要用来对数据库的 **数据**进行一些操作。

    ```mysql
    SELECT [列名] FROM 表名 [WHERE ....] --- 
    UPDATE 表名 SET [列名 = xx] [WHERE 列名 = *] ---
    INSERT INTO 表名 (列名1, 列名2) values (值1, 值2);
    DELETE FROM 表名 [WHERE 列名= **] --
    ```

* DDL

  * DDL (data definition language) 数据库定义语言

    * 比如说我们创建数据库用户的sql 语句, `CREATE`, `ALTER`, `DROP` 等, DDL 主要就是用在定义或者修改表的结构, 数据的类型, 表之间的连接和约束等初始化工作之上.

      ```mysql
      CREATE TABLE 表名(
          列名称 数据类型, 
          列名称, 数据类类型,
          ...
      )engine=InnoDB;
      
      ALTER TABLE table_name;
      ALTER COLUMN column_name datetype;
      
      DROP TABLE 表名称;
      DROP DATABASE 数据库的名称;
      ```

* DCL 

  * DCL (data control language) 数据库的控制语言
    * 用来设置或者修改数据库用户,或者角色权限的语句, 包括 `grant`, `deny`, `revoke` 等