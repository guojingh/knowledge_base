# 练习2-3：个性化消息 　用变量表示一个人的名字，并向其显示一条消息。显示的消息应非常简单.
# Hello Eric, would you like to learn some Python today?
name = "Eric"
print(f"{name}, would you like to learn some Python today?")

# 调整名字的大小写 　用变量表示一个人的名字，再以小写、大写和首字母大写的方式显示这个人名
name = "Eric"
# 小写
print(name.lower())
# 大写
print(name.upper())
# 首字母大写
print(name.title())

# 练习2-5：名言 　找一句你钦佩的名人说的名言，将其姓名和名言打印出来。输出应类似于下面这样（包括引号）
# Albert Einstein once said, “A person who never made a mistake never tried anything new.
print("Albert Einstein once said,“A person who never made a mistake never tried anything new.”")

# 练习2-6：名言2 　重复练习2-5，但用变量famous_person 表示名人的姓名，再创建要显示的消息并将其赋给变量message ，
# 然后打印这条消息。
famous_person = "Albert Einstein"
message = f"{famous_person} once said,“A person who never made a mistake never tried anything new.”"
print(message)

# 练习2-7：剔除人名中的空白 　用变量表示一个人的名字，并在其开头和末尾都包含一些空白字符。
# 务必至少使用字符组合"\t" 和"\n" 各一次。
# 打印这个人名，显示其开头和末尾的空白。然后，分别使用剔除函数lstrip() 、rstrip() 和strip() 对人名进行处理，
# 并将结果打印出来。
name = "\tZh\ne\tbiaoxian\n"
print(name)
print(name.lstrip())
print(name.rstrip())
print(name.strip())