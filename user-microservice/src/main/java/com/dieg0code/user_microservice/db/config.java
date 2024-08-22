package com.dieg0code.user_microservice.db;

import org.springframework.boot.jdbc.DataSourceBuilder;
import org.springframework.context.annotation.Bean;

import javax.sql.DataSource;

public class config {

    public static final String URL = "jdbc:postgresql://localhost:5432/userdb";
    public static final String USER = "test";
    public static final String PASSWORD = "test";

    @Bean
    public DataSource getDataSource() {
        DataSourceBuilder dataSourceBuilder = DataSourceBuilder.create();
        dataSourceBuilder.driverClassName("org.postgresql.Driver");
        dataSourceBuilder.url(URL);
        dataSourceBuilder.username(USER);
        dataSourceBuilder.password(PASSWORD);
        return dataSourceBuilder.build();
    }
}
