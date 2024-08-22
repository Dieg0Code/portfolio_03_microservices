package com.dieg0code.user_microservice.security;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.web.authentication.WebAuthenticationDetailsSource;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.util.List;

public class JwtAuthenticationFilter extends OncePerRequestFilter {

    @Value("${JWT_SECRET_KEY}")
    private String jwtSecretKey;

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {
        String header = request.getHeader("Authorization");

        if (header == null || !header.startsWith("Bearer ")) {
            filterChain.doFilter(request, response);
            return;
        }

        String token = header.replace("Bearer ", "");
        Claims claims = Jwts.parserBuilder()
                .setSigningKey(jwtSecretKey.getBytes())
                .build()
                .parseClaimsJws(token)
                .getBody();

        String username = claims.getSubject();
        request.setAttribute("username", username);
        int userID = claims.get("userID", Integer.class);
        String role = claims.get("role", String.class);
        List<GrantedAuthority> authorities = List.of(new SimpleGrantedAuthority("ROLE_" + role));

        UsernamePasswordAuthenticationToken authentication = new UsernamePasswordAuthenticationToken(
                userID, null, authorities
        );

        authentication.setDetails(new WebAuthenticationDetailsSource().buildDetails(request));
        SecurityContextHolder.getContext().setAuthentication(authentication);

        filterChain.doFilter(request, response);

    }
}
