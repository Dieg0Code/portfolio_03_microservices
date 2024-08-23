package com.dieg0code.user_microservice;

import com.dieg0code.user_microservice.json.request.CreateUserRequest;
import com.dieg0code.user_microservice.json.request.LoginRequest;
import com.dieg0code.user_microservice.json.response.UserResponse;
import com.dieg0code.user_microservice.models.User;
import com.dieg0code.user_microservice.repository.UserRepository;
import com.dieg0code.user_microservice.service.UserServiceImpl;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.test.util.ReflectionTestUtils;

import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.Mockito.*;

public class UserServiceImplTest {

    @Mock
    private UserRepository userRepository;

    @Mock
    private PasswordEncoder passwordEncoder;

    @InjectMocks
    private UserServiceImpl userService;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
        ReflectionTestUtils.setField(userService, "jwtSecretKey", "testSecretKey12345678901234567890");
    }

    @Test
    public void testCreateUserSuccess() {
        CreateUserRequest createUserRequest = new CreateUserRequest();
        createUserRequest.setUsername("test");
        createUserRequest.setPassword("test");
        createUserRequest.setEmail("test@test.com");
        createUserRequest.setRole("USER");

        User user = new User();

        user.setUserID(1);
        user.setUsername("test");
        user.setPassword("test");
        user.setEmail("test@test.com");
        user.setRole("USER");

        when(passwordEncoder.encode(createUserRequest.getPassword())).thenReturn("encodedPassword");
        when(userRepository.save(any(User.class))).thenReturn(user);

        int userID = userService.createUser(createUserRequest);

        assertEquals(1, userID);
        assertEquals("test", user.getUsername());
        assertEquals("test", user.getPassword());
        assertEquals("test@test.com", user.getEmail());
        assertEquals("USER", user.getRole());

        verify(passwordEncoder).encode(createUserRequest.getPassword());
        verify(userRepository, times(1)).save(any(User.class));
    }

    @Test
    public void testCreateUserFailure() {
        CreateUserRequest createUserRequest = new CreateUserRequest();
        createUserRequest.setUsername("test");
        createUserRequest.setPassword("test");
        createUserRequest.setEmail("test@test.com");
        createUserRequest.setRole("USER");

        when(passwordEncoder.encode(createUserRequest.getPassword())).thenThrow(new RuntimeException("Error"));

        int userID = userService.createUser(createUserRequest);

        assertEquals(-1, userID);

        verify(passwordEncoder).encode(createUserRequest.getPassword());
        verify(userRepository, never()).save(any(User.class));

    }

    @Test
    public void testGetUserSuccess() {
        User user = new User();

        user.setUserID(1);
        user.setUsername("test");
        user.setPassword("test");
        user.setEmail("test@test.com");
        user.setRole("USER");

        when(userRepository.findById(1)).thenReturn(Optional.of(user));

        UserResponse userResponse = userService.getUser(1);

        assertNotNull(userResponse);
        assertEquals(1, userResponse.getUserID());
        assertEquals("test", userResponse.getUsername());
        assertEquals("test@test.com", userResponse.getEmail());
        assertEquals("USER", userResponse.getRole());

        verify(userRepository, times(1)).findById(1);
    }

    @Test
    public void testGetUserFailure() {
        when(userRepository.findById(1)).thenReturn(Optional.empty());

        UserResponse userResponse = userService.getUser(1);

        assertNull(userResponse);

        verify(userRepository, times(1)).findById(1);
    }

    @Test
    public void testUpdateUserSuccess() {
        CreateUserRequest createUserRequest = new CreateUserRequest();
        createUserRequest.setUsername("test");
        createUserRequest.setPassword("test");
        createUserRequest.setEmail("test@test.com");
        createUserRequest.setRole("USER");

        User user = new User();

        user.setUserID(1);
        user.setUsername("old");
        user.setPassword("old");
        user.setEmail("old@email.com");
        user.setRole("USER");

        when(userRepository.findById(1)).thenReturn(Optional.of(user));
        when(passwordEncoder.encode(createUserRequest.getPassword())).thenReturn("encodedPassword");
        when(userRepository.save(any(User.class))).thenReturn(user);

        boolean updated = userService.updateUser(1, createUserRequest);

        assertTrue(updated);
        assertEquals("test", user.getUsername());
        assertEquals("encodedPassword", user.getPassword());
        assertEquals("test@test.com", user.getEmail());

        verify(userRepository, times(1)).findById(1);
        verify(passwordEncoder).encode(createUserRequest.getPassword());
        verify(userRepository, times(1)).save(any(User.class));
    }

    @Test
    public void testUpdateUserFailure() {
        CreateUserRequest createUserRequest = new CreateUserRequest();
        createUserRequest.setUsername("test");
        createUserRequest.setPassword("test");
        createUserRequest.setEmail("test@test.com");
        createUserRequest.setRole("USER");

        when(userRepository.findById(1)).thenReturn(Optional.empty());

        boolean updated = userService.updateUser(1, createUserRequest);

        assertFalse(updated);

        verify(userRepository, times(1)).findById(1);
        verify(passwordEncoder, never()).encode(createUserRequest.getPassword());
        verify(userRepository, never()).save(any(User.class));
    }

    @Test
    public void testDeleteUserSuccess() {
        User user = new User();
        user.setUserID(1);

        when(userRepository.findById(1)).thenReturn(Optional.of(user));

        boolean deleted = userService.deleteUser(1);

        assertTrue(deleted);
        verify(userRepository, times(1)).deleteById(1);
    }

    @Test
    public void testDeleteUserFailure() {
        when(userRepository.findById(1)).thenReturn(Optional.empty());

        boolean deleted = userService.deleteUser(1);

        assertFalse(deleted);
        verify(userRepository, never()).deleteById(1);
    }

    @Test
    void testLoginSuccessful() {
        // Arrange
        LoginRequest loginRequest = new LoginRequest();
        loginRequest.setEmail("test@test.com");
        loginRequest.setPassword("password123");

        User user = new User();
        user.setEmail("test@test.com");
        user.setPassword("encodedPassword");
        user.setUsername("testUser");
        user.setUserID(1);
        user.setRole("USER");

        when(userRepository.findByEmail("test@test.com")).thenReturn(Optional.of(user));
        when(passwordEncoder.matches("password123", "encodedPassword")).thenReturn(true);

        // Act
        String token = userService.login(loginRequest);

        // Assert
        assertNotNull(token, "Token should not be null");
        assertTrue(token.startsWith("Bearer "), "Token should start with 'Bearer '");

        // Verify
        verify(userRepository).findByEmail("test@test.com");
        verify(passwordEncoder).matches("password123", "encodedPassword");
    }

    @Test
    void testLoginUserNotFound() {
        // Arrange
        LoginRequest loginRequest = new LoginRequest();
        loginRequest.setEmail("nonexistent@example.com");
        loginRequest.setPassword("password123");

        when(userRepository.findByEmail("nonexistent@example.com")).thenReturn(Optional.empty());

        // Act
        String token = userService.login(loginRequest);

        // Assert
        assertNull(token, "Token should be null for non-existent user");

        // Verify
        verify(userRepository).findByEmail("nonexistent@example.com");
        verifyNoInteractions(passwordEncoder);
    }

    @Test
    void testLoginInvalidPassword() {
        // Arrange
        LoginRequest loginRequest = new LoginRequest();
        loginRequest.setEmail("test@test.com");
        loginRequest.setPassword("wrongPassword");

        User user = new User();
        user.setEmail("test@test.com");
        user.setPassword("encodedPassword");

        when(userRepository.findByEmail("test@test.com")).thenReturn(Optional.of(user));
        when(passwordEncoder.matches("wrongPassword", "encodedPassword")).thenReturn(false);

        // Act
        String token = userService.login(loginRequest);

        // Assert
        assertNull(token, "Token should be null for invalid password");

        // Verify
        verify(userRepository).findByEmail("test@test.com");
        verify(passwordEncoder).matches("wrongPassword", "encodedPassword");
    }

    @Test
    void testLoginExceptionHandling() {
        // Arrange
        LoginRequest loginRequest = new LoginRequest();
        loginRequest.setEmail("test@test.com");
        loginRequest.setPassword("password123");

        when(userRepository.findByEmail("test@test.com")).thenThrow(new RuntimeException("Database error"));

        // Act
        String token = userService.login(loginRequest);

        // Assert
        assertNull(token, "Token should be null when an exception occurs");

        // Verify
        verify(userRepository).findByEmail("test@test.com");
        verifyNoInteractions(passwordEncoder);
    }

    @Test
    void testLoginEmptyCredentials() {
        // Arrange
        LoginRequest loginRequest = new LoginRequest();
        loginRequest.setEmail("");
        loginRequest.setPassword("");

        // Act
        String token = userService.login(loginRequest);

        // Assert
        assertNull(token, "Token should be null for empty credentials");

        // Verify
        verifyNoInteractions(userRepository);
        verifyNoInteractions(passwordEncoder);
    }


}
