package org.apache.dubbo;

import java.io.Serializable;
import java.util.List;

//pojo Page
public class Page<T> implements Serializable {
    private static final long serialVersionUID = 1L;
    private List<T> data;

    public Page() {
    }

    public Page(List<T> data) {
        this.data = data;
    }

    public List<T> getData() {
        return this.data;
    }

    public void setData(List<T> data) {
        this.data = data;
    }
}