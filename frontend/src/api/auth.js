import axios from "axios";
import {useState} from "react";

let accessToken = JSON.parse(localStorage.getItem('ACCESS_TOKEN')) || null

export const getExpirationDate = (jwtToken) => {
    if (!jwtToken) {
        return null;
    }

    const jwt = JSON.parse(atob(jwtToken.split('.')[1]));

    return jwt && jwt.exp && jwt.exp * 1000 || null;
};

export const isExpired = (exp) => {
    if (!exp) {
        return false;
    }

    return Date.now() > exp;
};

export const getToken = () => {
    if (!accessToken) {
        return null;
    }

    if (isExpired(getExpirationDate(accessToken))) {
        refreshToken();
    }

    return accessToken;
};

export const refreshToken = () => {
    let updatedToken, refreshError

    axios({
        method: 'post',
        url: '/api/v1/users/auth/refresh',
        withCredentials: true,
        data: {
            fingerprint: "fingerprint"
        }
    })
        .then(resp => {
            updatedToken = resp.data.data;
            setToken(updatedToken);
        })
        .catch(error => {
            console.log(error.response.data);
            refreshError = error.response.data;
        })
}

export const signIn = (event) => {
    event.preventDefault();

    if (isLoggedIn())
        return

    const data = new FormData(event.currentTarget);
    let token = null

    axios({
        method: 'post',
        url: '/api/v1/users/auth/sign-in',
        withCredentials: true,
        data: {
            email: data.get('email'),
            password: data.get('password'),
            fingerprint: "fingerprint"
        }
    })
        .then(resp => {
            token = resp.data.data;
            setToken(token);
            console.log(token);
            window.location = '/catalog';
        })
        .catch(error => {
            console.log(error.response.data);
        })
};

export const signUp = (event) => {
    event.preventDefault();

    const data = new FormData(event.currentTarget);

    return axios({
        method: 'post',
        url: '/api/v1/users/auth/sign-up',
        withCredentials: true,
        data: {
            email: data.get('email'),
            password: data.get('password'),
            name: data.get('name')
        }
    })
};

export const setToken = (token) => {
    if (token) {
        localStorage.setItem('ACCESS_TOKEN', JSON.stringify(token));
    } else {
        localStorage.removeItem('ACCESS_TOKEN');
    }

    accessToken = token;
};

export const isLoggedIn = () => {
    return !!accessToken;
};
