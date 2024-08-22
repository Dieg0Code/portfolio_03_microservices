package com.dieg0code.user_microservice.json.response;

import lombok.Data;

@Data
public class UserResponse {
    private int userID;
    private String username;
    private String email;
    private String role;
}
