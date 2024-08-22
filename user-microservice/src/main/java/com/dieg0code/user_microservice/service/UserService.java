package com.dieg0code.user_microservice.service;

import com.dieg0code.user_microservice.json.request.CreateUserRequest;
import com.dieg0code.user_microservice.json.response.UserResponse;

import java.util.List;

public interface UserService {
    int createUser(CreateUserRequest createUserRequest);
    UserResponse getUser(int userID);
    boolean updateUser(int userID, CreateUserRequest createUserRequest);
    boolean deleteUser(int userID);
    List<UserResponse> getAllUsers();

}
