package org.apache.dubbo;

import java.io.Serializable;

//pojo UserVo
public class UserVo implements Serializable {
    private User user;
    public UserVo(User user) { this.user = user; }
    public UserVo() { }
    public User getUser() {
        return user;
    }
    public void setUser(User user) {
        this.user = user;
    }
}