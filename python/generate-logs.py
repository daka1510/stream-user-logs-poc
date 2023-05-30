import logging
import time

def get_user_message_logger():
    log = logging.getLogger("user-message-poc")
    file_handler = logging.FileHandler('../shared-fs/user-logs.log') 
    log.addHandler(file_handler)
    return log
    
user_message_logger = get_user_message_logger()

time.sleep(10)

for number in range(120):
    time.sleep(0.3)
    user_message_logger.fatal(f"log statement {number}")
    
