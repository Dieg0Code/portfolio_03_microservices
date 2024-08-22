package com.dieg0code.user_microservice.repository;

import com.dieg0code.user_microservice.Models.User;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface UserRepository extends JpaRepository<User, Integer> {}
