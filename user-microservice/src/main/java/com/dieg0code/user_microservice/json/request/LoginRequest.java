package com.dieg0code.user_microservice.json.request;

import lombok.Data;

@Data
public class LoginRequest {
    private String email;
    private String password;
}
