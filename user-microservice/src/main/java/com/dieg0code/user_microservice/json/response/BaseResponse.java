package com.dieg0code.user_microservice.json.response;

import lombok.Data;

@Data
public class BaseResponse<T> {
    private String code;
    private String status;
    private String msg;
    private T data;
}
