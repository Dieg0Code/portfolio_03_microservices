package com.dieg0code.user_microservice.handlers;

import com.dieg0code.user_microservice.json.response.BaseResponse;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.AccessDeniedException;
import org.springframework.security.core.AuthenticationException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;

@ControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(AccessDeniedException.class)
    public ResponseEntity<BaseResponse<String>> handleAccessDeniedException(AccessDeniedException ex) {
        BaseResponse<String> response = new BaseResponse<>();
        response.setCode("403");
        response.setStatus("error");
        response.setMsg("Access Denied: " + ex.getMessage());
        response.setData(null);
        return new ResponseEntity<>(response, HttpStatus.FORBIDDEN);
    }

    @ExceptionHandler(AuthenticationException.class)
    public ResponseEntity<BaseResponse<String>> handleAuthenticationException(AuthenticationException ex) {
        BaseResponse<String> response = new BaseResponse<>();
        response.setCode("401");
        response.setStatus("error");
        response.setMsg("Unauthorized: " + ex.getMessage());
        response.setData(null);
        return new ResponseEntity<>(response, HttpStatus.UNAUTHORIZED);
    }

}

