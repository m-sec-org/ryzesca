#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    :  2023/8/16 9:33
# @Author  :  buding
# @Email   :  budingplus@163.com
# @File    :  sql_insert_use.py
# @software:  PyCharm
import sqlite3
import time

import pymysql


# 查询数据库总数的sql
#

def initMysqlClient():
    conn = pymysql.connect(host="127.0.0.1", port=3306, user="root", passwd="123456", db="themis", charset="utf8")
    return conn
    pass


def initSqliteClient():
    conn = sqlite3.connect("themis.db")
    return conn
    pass


f = open('themis_error.sql', 'a', encoding='utf-8')


def migrate():
    mconn = initMysqlClient()
    sconn = initSqliteClient()
    mursor = mconn.cursor()
    # 查询出sqlite所有的表
    cursor = sconn.cursor()
    cursor.execute("select name from sqlite_master where type='table' order by name")
    tables = cursor.fetchall()
    start_time = time.time()
    sql = ''
    aa = 0
    all_data = 0
    for table in tables:
        if table[0] == 'cve_infos':
            print("开始插入表%s" % table[0])
            # 查询出每个表的所有数据
            cursor.execute("select * from %s" % table)
            datas = cursor.fetchall()
            # 使用pymysql防止sql注入的方式拼接处插入mysql的语句
            for data in datas:
                sql = "insert into %s values " % table
                sql += "("
                for i in range(len(data)):
                    if i == 0:
                        sql += "%s"
                    else:
                        sql += ",%s"
                sql += ")"
                try:
                    mursor.execute(sql, data)
                    mconn.commit()
                    all_data += 1
                    # print("这是第%s条数据" % all_data)
                except Exception as e:
                    print("插入失败，失败原因：%s" % e)
                    f.write(sql % data + "\n")
                    pass
                pass
    end_time = time.time()
    print("插入完成，共插入%s条数据，耗时%s秒" % (all_data, end_time - start_time))
    mconn.close()
    sconn.close()


if __name__ == '__main__':
    migrate()
    pass
