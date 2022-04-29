import React from 'react';
import {Alert, AlertTitle} from "@mui/material";

import './Error.css'

const Error = ({message}) => {
    return (
        <div className='error-container'>
            <Alert severity="error">
                <AlertTitle>Error</AlertTitle>
                {message}
            </Alert>
        </div>
    );
};

export default Error;