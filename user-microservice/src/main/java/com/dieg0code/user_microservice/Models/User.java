package com.dieg0code.user_microservice.models;

import jakarta.persistence.*;
import lombok.Data;

@Data
@Entity
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private int userID;

    @Column
    private String username;

    @Column
    private String password;

    @Column
    private String email;

    @Column
    private String role;
}
