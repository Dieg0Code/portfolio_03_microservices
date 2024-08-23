package com.dieg0code.user_microservice;

import com.dieg0code.user_microservice.controller.UserControllerImpl;
import com.dieg0code.user_microservice.json.request.CreateUserRequest;
import com.dieg0code.user_microservice.json.request.LoginRequest;
import com.dieg0code.user_microservice.json.response.UserResponse;
import com.dieg0code.user_microservice.service.UserService;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.setup.MockMvcBuilders;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.util.ArrayList;
import java.util.List;

import static org.mockito.Mockito.*;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.*;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;


public class UserControllerImplTest {

    private MockMvc mockMvc;

    @Mock
    private UserService userService;

    @InjectMocks
    private UserControllerImpl userController;

    private ObjectMapper objectMapper = new ObjectMapper();

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
        mockMvc = MockMvcBuilders.standaloneSetup(userController).build();
    }

    @Test
    void testCreateUserSuccess() throws Exception {
        CreateUserRequest request = new CreateUserRequest();
        request.setUsername("test");
        request.setPassword("test");
        request.setEmail("test@test.com");
        request.setRole("USER");

        when(userService.createUser(any(CreateUserRequest.class))).thenReturn(1);

        mockMvc.perform(post("/user/create")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(request)))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("200"))
                .andExpect(jsonPath("$.status").value("success"))
                .andExpect(jsonPath("$.data").value(1));

        verify(userService).createUser(any(CreateUserRequest.class));
    }

    @Test
    void testCreateUserFailure() throws Exception {
        CreateUserRequest request = new CreateUserRequest();
        request.setUsername("test");
        request.setPassword("test");
        request.setEmail("test@test.com");
        request.setRole("USER");


        when(userService.createUser(any(CreateUserRequest.class))).thenReturn(-1);

        mockMvc.perform(post("/user/create")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(request)))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("500"))
                .andExpect(jsonPath("$.status").value("error"))
                .andExpect(jsonPath("$.data").doesNotExist());

        verify(userService).createUser(any(CreateUserRequest.class));
    }

    @Test
    void testGetUserSuccess() throws Exception {
        int userId = 1;
        UserResponse userResponse = new UserResponse();
        userResponse.setUserID(userId);
        userResponse.setUsername("test");
        userResponse.setEmail("test@test.com");

        when(userService.getUser(userId)).thenReturn(userResponse);

        mockMvc.perform(get("/user/{userId}", userId))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("200"))
                .andExpect(jsonPath("$.status").value("success"));

        verify(userService).getUser(userId);
    }

    @Test
    void testGetUserNotFound() throws Exception {
        int userId = 1; // invalid user id

        when(userService.getUser(userId)).thenReturn(null);

        mockMvc.perform(get("/user/{userId}", userId))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("404"))
                .andExpect(jsonPath("$.status").value("error"))
                .andExpect(jsonPath("$.data").doesNotExist());

        verify(userService).getUser(userId);
    }

    @Test
    void testGetUserInternalError() throws Exception {
        int userId = 1;

        when(userService.getUser(userId)).thenThrow(new RuntimeException("Error"));

        mockMvc.perform(get("/user/{userId}", userId))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("500"))
                .andExpect(jsonPath("$.status").value("error"))
                .andExpect(jsonPath("$.data").doesNotExist());

        verify(userService).getUser(userId);
    }

    @Test
    void testUpdateUserSuccess() throws Exception {
        int userId = 1;
        CreateUserRequest request = new CreateUserRequest();
        request.setUsername("test");
        request.setPassword("test");
        request.setEmail("test@test.com");
        request.setRole("USER");

        when(userService.updateUser(eq(userId), any(CreateUserRequest.class))).thenReturn(true);

        mockMvc.perform(put("/user/{userId}", userId)
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(request)))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("200"))
                .andExpect(jsonPath("$.status").value("success"))
                .andExpect(jsonPath("$.data").value(true));

        verify(userService).updateUser(eq(userId), any(CreateUserRequest.class));
    }

    @Test
    void testUpdateUserFailure() throws Exception {
        int userId = 1;
        CreateUserRequest request = new CreateUserRequest();
        request.setUsername("test");
        request.setPassword("test");
        request.setEmail("test@test.com");
        request.setRole("USER");

        when(userService.updateUser(eq(userId), any(CreateUserRequest.class))).thenReturn(false);

        mockMvc.perform(put("/user/{userId}", userId)
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(request)))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("500"))
                .andExpect(jsonPath("$.status").value("error"))
                .andExpect(jsonPath("$.data").value(false));

        verify(userService).updateUser(eq(userId), any(CreateUserRequest.class));
    }

    @Test
    void testUpdateUserIternalError() throws Exception {
        int userId = 1;
        CreateUserRequest request = new CreateUserRequest();
        request.setUsername("test");
        request.setPassword("test");
        request.setEmail("test@test.com");
        request.setRole("USER");

        when(userService.updateUser(eq(userId), any(CreateUserRequest.class))).thenThrow(new RuntimeException("Error"));

        mockMvc.perform(put("/user/{userId}", userId)
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(request)))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("500"))
                .andExpect(jsonPath("$.status").value("error"))
                .andExpect(jsonPath("$.data").doesNotExist());

        verify(userService).updateUser(eq(userId), any(CreateUserRequest.class));

    }

    @Test
    void testDeleteUserSuccess() throws Exception {
        int userId = 1;

        when(userService.deleteUser(userId)).thenReturn(true);

        mockMvc.perform(delete("/user/{userId}", userId))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("200"))
                .andExpect(jsonPath("$.status").value("success"))
                .andExpect(jsonPath("$.data").value(true));

        verify(userService).deleteUser(userId);
    }

    @Test
    void testDeleteUserFailure() throws Exception {
            int userId = 1;

        when(userService.deleteUser(userId)).thenReturn(false);

        mockMvc.perform(delete("/user/{userId}", userId))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("404"))
                .andExpect(jsonPath("$.status").value("error"))
                .andExpect(jsonPath("$.data").value(false));

        verify(userService).deleteUser(userId);
    }

    @Test
    void testDeleteUserInternalError() throws Exception {
        int userId = -1; // invalid id

        when(userService.deleteUser(userId)).thenThrow(new RuntimeException("Error"));

        mockMvc.perform(delete("/user/{userId}", userId))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("500"))
                .andExpect(jsonPath("$.status").value("error"))
                .andExpect(jsonPath("$.data").doesNotExist());

        verify(userService).deleteUser(userId);
    }

    @Test
    void testGetAllUsersSuccess() throws Exception {
        List<UserResponse> userResponses = new ArrayList<>();

        UserResponse userOne = new UserResponse();
        userOne.setUserID(1);
        userOne.setUsername("test");
        userOne.setEmail("one@test.com");

        userResponses.add(userOne);

        UserResponse userTwo = new UserResponse();
        userTwo.setUserID(2);
        userTwo.setUsername("test2");
        userTwo.setEmail("two@test.com");

        userResponses.add(userTwo);

        when(userService.getAllUsers()).thenReturn(userResponses);

        mockMvc.perform(get("/user/all"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("200"))
                .andExpect(jsonPath("$.status").value("success"));

        verify(userService).getAllUsers();
    }

    @Test
    void testGetAllUsersFailure() throws Exception {

        when(userService.getAllUsers()).thenReturn(null);

        mockMvc.perform(get("/user/all"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("404"))
                .andExpect(jsonPath("$.status").value("error"))
                .andExpect(jsonPath("$.data").doesNotExist());

        verify(userService).getAllUsers();

    }

    @Test
    void testGetAllUsersInternalError() throws Exception {

        when(userService.getAllUsers()).thenThrow(new RuntimeException("Error"));

        mockMvc.perform(get("/user/all"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("500"))
                .andExpect(jsonPath("$.status").value("error"))
                .andExpect(jsonPath("$.data").doesNotExist());

        verify(userService).getAllUsers();

    }

    @Test
    void testLogin() throws Exception {
        LoginRequest loginRequest = new LoginRequest();
        loginRequest.setEmail("test@tes.com");
        loginRequest.setPassword("test");

        String token = "sample_token";
        when(userService.login(any(LoginRequest.class))).thenReturn(token);

        mockMvc.perform(post("/user/login")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(loginRequest)))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("200"))
                .andExpect(jsonPath("$.status").value("success"))
                .andExpect(jsonPath("$.data.token").value(token));

        verify(userService).login(any(LoginRequest.class));
    }

    @Test
    void testLoginFailure() throws Exception {
        LoginRequest loginRequest = new LoginRequest();
        loginRequest.setEmail("test@test.com");

        when(userService.login(any(LoginRequest.class))).thenReturn(null);

        mockMvc.perform(post("/user/login")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(loginRequest)))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("401"))
                .andExpect(jsonPath("$.status").value("error"))
                .andExpect(jsonPath("$.data").doesNotExist());

        verify(userService).login(any(LoginRequest.class));
    }

    @Test
    void testLoginInternalError() throws Exception {
        LoginRequest loginRequest = new LoginRequest();
        loginRequest.setEmail("test@test.com");
        loginRequest.setPassword("test");


        when(userService.login(any(LoginRequest.class))).thenThrow(new RuntimeException("Error"));

        mockMvc.perform(post("/user/login")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content(objectMapper.writeValueAsString(loginRequest)))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.code").value("500"))
                .andExpect(jsonPath("$.status").value("error"))
                .andExpect(jsonPath("$.data").doesNotExist());

        verify(userService).login(any(LoginRequest.class));
    }
}
