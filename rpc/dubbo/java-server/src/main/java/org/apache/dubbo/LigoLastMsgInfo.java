package org.apache.dubbo;

import java.io.Serializable;
import java.util.*;

public class LigoLastMsgInfo implements Serializable {

    private int messageId;

    private String text;

    private Date messageTime = new Date();

}
