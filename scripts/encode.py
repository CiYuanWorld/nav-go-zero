import base64

key_data = "{0}{1}".format("hello", "悟2空")

print(base64.b64encode(key_data.encode()))



# key_data_bytes = bytes(key_data, "UTF-8")
# key_data_bytes = key_data.encode()
# https://blog.csdn.net/faihung/article/details/90516180
