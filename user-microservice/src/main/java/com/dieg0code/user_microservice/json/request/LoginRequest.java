package com.dieg0code.user_microservice.json.request;

import lombok.Data;

@Data
public class LoginRequest {
    private String username;
    private String password;
}
