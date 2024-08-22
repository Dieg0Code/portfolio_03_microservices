package com.dieg0code.user_microservice.controller;

import com.dieg0code.user_microservice.json.request.CreateUserRequest;
import com.dieg0code.user_microservice.json.request.LoginRequest;
import com.dieg0code.user_microservice.json.response.BaseResponse;
import com.dieg0code.user_microservice.json.response.LoginResponse;
import com.dieg0code.user_microservice.json.response.UserResponse;
import com.dieg0code.user_microservice.service.UserService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Slf4j
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
                log.info("User created with ID: {}", userId);
            } else {
                response.setCode("500");
                response.setStatus("error");
                response.setMsg("Error creating user");
                response.setData(null);
                log.error("Error creating user");
            }
        } catch (Exception e) {
            response.setCode("500");
            response.setStatus("error");
            response.setMsg("Exception occurred: " + e.getMessage());
            response.setData(null);
            log.error("Error creating user: {}", e.getMessage());
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
                log.info("User retrieved with ID: {}", userID);
            } else {
                response.setCode("404");
                response.setStatus("error");
                response.setMsg("User not found");
                response.setData(null);
                log.error("User with ID: {} not found", userID);
            }
        } catch (Exception e) {
            response.setCode("500");
            response.setStatus("error");
            response.setMsg("Exception occurred: " + e.getMessage());
            response.setData(null);
            log.error("Error getting user with ID: {}", userID, e);
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
                log.error("Error updating user with ID: {}", userID);
            }
        } catch (Exception e) {
            response.setCode("500");
            response.setStatus("error");
            response.setMsg("Exception occurred: " + e.getMessage());
            response.setData(false);
            log.error("Error updating user with ID: {}", userID, e);
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
                log.info("User deleted with ID: {}", userID);
            } else {
                response.setCode("500");
                response.setStatus("error");
                response.setMsg("Error deleting user");
                response.setData(false);
                log.error("Error deleting user with ID: {}", userID);
            }
        } catch (Exception e) {
            response.setCode("500");
            response.setStatus("error");
            response.setMsg("Exception occurred: " + e.getMessage());
            response.setData(false);
            log.error("Error deleting user with ID: {}", userID, e);
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
                log.info("All users retrieved");
            } else {
                response.setCode("404");
                response.setStatus("error");
                response.setMsg("No users found");
                response.setData(null);
                log.error("No users found");
            }
        } catch (Exception e) {
            response.setCode("500");
            response.setStatus("error");
            response.setMsg("Exception occurred: " + e.getMessage());
            response.setData(null);
            log.error("Error getting all users: {}", e.getMessage());
        }

        return response;
    }

    @PostMapping("/login")
    @Override
    public BaseResponse<LoginResponse> login(@RequestBody LoginRequest loginRequest) {
        BaseResponse<LoginResponse> response = new BaseResponse<>();

         try {
                String token = userService.login(loginRequest);

                if (token != null) {
                 LoginResponse loginResponse = new LoginResponse();
                 loginResponse.setToken(token);

                 response.setCode("200");
                 response.setStatus("success");
                 response.setMsg("Login successful");
                 response.setData(loginResponse);
                 log.info("Login successful");
                } else {
                 response.setCode("401");
                 response.setStatus("error");
                 response.setMsg("Invalid credentials");
                 response.setData(null);
                 log.error("Invalid credentials");
                }
          } catch (Exception e) {
                response.setCode("500");
                response.setStatus("error");
                response.setMsg("Exception occurred: " + e.getMessage());
                response.setData(null);
                log.error("Error logging in: {}", e.getMessage());
          }

            return response;
    }
}
