package com.dieg0code.user_microservice.controller;

import com.dieg0code.user_microservice.json.request.CreateUserRequest;
import com.dieg0code.user_microservice.json.response.BaseResponse;
import com.dieg0code.user_microservice.json.response.UserResponse;
import com.dieg0code.user_microservice.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/user")
public class UserControllerImpl implements UserController {

    @Autowired
    private UserService userService;

    @PostMapping("/create")
    @Override
    public BaseResponse<Integer> createUser(@RequestBody CreateUserRequest createUserRequest) {
        BaseResponse<Integer> response = new BaseResponse<>();
        try {
            int userId = userService.createUser(createUserRequest);

            if (userId > -1) {
                response.setCode("200");
                response.setStatus("success");
                response.setMsg("User created successfully");
                response.setData(userId);
            } else {
                response.setCode("500");
                response.setStatus("error");
                response.setMsg("Error creating user");
                response.setData(null);
            }
        } catch (Exception e) {
            response.setCode("500");
            response.setStatus("error");
            response.setMsg("Exception occurred: " + e.getMessage());
            response.setData(null);
        }

        return response;
    }

    @GetMapping("/{userID}")
    @Override
    public BaseResponse<UserResponse> getUser(@PathVariable int userID) {
        BaseResponse<UserResponse> response = new BaseResponse<>();

        try {
            UserResponse userResponse = userService.getUser(userID);

            if (userResponse != null) {
                response.setCode("200");
                response.setStatus("success");
                response.setMsg("User retrieved successfully");
                response.setData(userResponse);
            } else {
                response.setCode("404");
                response.setStatus("error");
                response.setMsg("User not found");
                response.setData(null);
            }
        } catch (Exception e) {
            response.setCode("500");
            response.setStatus("error");
            response.setMsg("Exception occurred: " + e.getMessage());
            response.setData(null);
        }

        return response;
    }

    @PutMapping("/{userID}")
    @Override
    public BaseResponse<Boolean> updateUser(@PathVariable int userID, @RequestBody CreateUserRequest createUserRequest) {
        BaseResponse<Boolean> response = new BaseResponse<>();

        try {
            boolean updated = userService.updateUser(userID, createUserRequest);

            if (updated) {
                response.setCode("200");
                response.setStatus("success");
                response.setMsg("User updated successfully");
                response.setData(true);
            } else {
                response.setCode("500");
                response.setStatus("error");
                response.setMsg("Error updating user");
                response.setData(false);
            }
        } catch (Exception e) {
            response.setCode("500");
            response.setStatus("error");
            response.setMsg("Exception occurred: " + e.getMessage());
            response.setData(false);
        }

        return response;
    }

    @DeleteMapping("/{userID}")
    @Override
    public BaseResponse<Boolean> deleteUser(@PathVariable int userID) {
        BaseResponse<Boolean> response = new BaseResponse<>();

        try {
            boolean deleted = userService.deleteUser(userID);

            if (deleted) {
                response.setCode("200");
                response.setStatus("success");
                response.setMsg("User deleted successfully");
                response.setData(true);
            } else {
                response.setCode("500");
                response.setStatus("error");
                response.setMsg("Error deleting user");
                response.setData(false);
            }
        } catch (Exception e) {
            response.setCode("500");
            response.setStatus("error");
            response.setMsg("Exception occurred: " + e.getMessage());
            response.setData(false);
        }

        return response;
    }

    @GetMapping("/all")
    @Override
    public BaseResponse<List<UserResponse>> getAllUsers() {
        BaseResponse<List<UserResponse>> response = new BaseResponse<>();

        try {
            List<UserResponse> userResponses = userService.getAllUsers();

            if (userResponses != null) {
                response.setCode("200");
                response.setStatus("success");
                response.setMsg("Users retrieved successfully");
                response.setData(userResponses);
            } else {
                response.setCode("404");
                response.setStatus("error");
                response.setMsg("No users found");
                response.setData(null);
            }
        } catch (Exception e) {
            response.setCode("500");
            response.setStatus("error");
            response.setMsg("Exception occurred: " + e.getMessage());
            response.setData(null);
        }

        return response;
    }
}
