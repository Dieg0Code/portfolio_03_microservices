package com.dieg0code.user_microservice.service;

import com.dieg0code.user_microservice.Models.User;
import com.dieg0code.user_microservice.json.request.CreateUserRequest;
import com.dieg0code.user_microservice.json.response.UserResponse;
import com.dieg0code.user_microservice.repository.UserRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

@Slf4j
@Service
public class UserServiceImpl implements UserService{

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private BCryptPasswordEncoder bCryptPasswordEncoder;

    @Override
    public int createUser(CreateUserRequest createUserRequest) {
        int userID = -1;

        try {
            User user = new User();
            user.setUsername(createUserRequest.getUsername());
            String encryptedPassword = bCryptPasswordEncoder.encode(createUserRequest.getPassword());
            user.setPassword(encryptedPassword);
            user.setEmail(createUserRequest.getEmail());
            user.setRole(createUserRequest.getRole());

            User savedUser = userRepository.save(user);

            userID = savedUser.getUserID();
            log.info("User created with ID: {}", userID);
        } catch (Exception e) {
            log.error("Error creating user: {}", e.getMessage());
        }

        return userID;
    }

    @Override
    public UserResponse getUser(int userID) {
        UserResponse userResponse = null;

        try {
            Optional<User> userOptional = userRepository.findById(userID);

            if (userOptional.isPresent()) {
                User user = userOptional.get();

                userResponse = new UserResponse();
                userResponse.setUserID(user.getUserID());
                userResponse.setUsername(user.getUsername());
                userResponse.setEmail(user.getEmail());
                userResponse.setRole(user.getRole());
            } else {
                log.error("User with ID: {} not found", userID);
            }
        } catch (Exception e) {
            log.error("Error getting user with ID: {}", userID, e);
        }

        return userResponse;
    }

    @Override
    public boolean updateUser(int userID, CreateUserRequest createUserRequest) {
        try {
            Optional<User> userOptional = userRepository.findById(userID);

            if (userOptional.isPresent()) {
                User user = userOptional.get();
                user.setUsername(createUserRequest.getUsername());
                String encryptedPassword = bCryptPasswordEncoder.encode(createUserRequest.getPassword());
                user.setPassword(encryptedPassword);
                user.setEmail(createUserRequest.getEmail());
                user.setRole(createUserRequest.getRole());

                userRepository.save(user);
                log.info("User with ID: {} updated", userID);
                return true;
            } else {
                log.error("User with ID: {} not found", userID);
                return false;
            }

        } catch (Exception e) {
            log.error("Error updating user with ID: {}", userID, e);
            return false;
        }
    }

    @Override
    public boolean deleteUser(int userID) {
        try {
            Optional<User> userOptional = userRepository.findById(userID);

            if (userOptional.isPresent()) {
                userRepository.deleteById(userID);
                log.info("User with ID: {} deleted", userID);
                return true;
            } else {
                log.error("User with ID: {} not found", userID);
                return false;
            }

        } catch (Exception e) {
            log.error("Error deleting user with ID: {}", userID, e);
            return false;
        }
    }

    @Override
    public List<UserResponse> getAllUsers() {
        List<UserResponse> userResponses = new ArrayList<>();

        try {
            List<User> users = userRepository.findAll();

            for (User user : users) {
                UserResponse userResponse = new UserResponse();
                userResponse.setUserID(user.getUserID());
                userResponse.setUsername(user.getUsername());
                userResponse.setEmail(user.getEmail());
                userResponse.setRole(user.getRole());

                userResponses.add(userResponse);
            }
        } catch (Exception e) {
            log.error("Error getting all users: {}", e.getMessage());
        }

        return userResponses;
    }
}
