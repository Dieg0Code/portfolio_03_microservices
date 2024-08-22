package com.dieg0code.user_microservice.controller;

import com.dieg0code.user_microservice.json.request.CreateUserRequest;
import com.dieg0code.user_microservice.json.request.LoginRequest;
import com.dieg0code.user_microservice.json.response.BaseResponse;
import com.dieg0code.user_microservice.json.response.LoginResponse;
import com.dieg0code.user_microservice.json.response.UserResponse;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;

import java.util.List;

public interface UserController {
    BaseResponse<Integer> createUser(@RequestBody CreateUserRequest createUserRequest);
    BaseResponse<UserResponse> getUser(@PathVariable int userID);
    BaseResponse<Boolean> updateUser(@PathVariable int userID, @RequestBody CreateUserRequest createUserRequest);
    BaseResponse<Boolean> deleteUser(@PathVariable int userID);
    BaseResponse<List<UserResponse>> getAllUsers();
    BaseResponse<LoginResponse> login(@RequestBody LoginRequest loginRequest);
}
