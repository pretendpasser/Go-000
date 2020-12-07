学习笔记

For qusetion:
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

For answer:
sql link:https://pkg.go.dev/database/sql

define:
var ErrNoRows = errors.New("sql: no rows in result set")

return:
ErrNoRows is returned by Scan when QueryRow doesn't return a row. 
In such a case, QueryRow returns a placeholder *Row value that defers this error until a Scan. 

从dao层对数据库进行操作，在这里获取数据或者err，在此层的error不建议直接进行默认值或者降级处理，直接将该error上抛的上一层。上面service层调用dao层的函数时候，如果下层抛上来一个error，可以通过wrapf对之前的error增加新的信息，最终层层上抛，最后在最上层处打印log或者将stack信息完全打印出来以用来追踪问题。