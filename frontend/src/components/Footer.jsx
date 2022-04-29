import React from 'react';
import {Grid} from "@mui/material";
import Typography from "@mui/material/Typography";

import './Footer.css'

const Footer = () => {
    const footers = [
        {
            title: 'Company',
            description: ['Team', 'History', 'Contact us', 'Locations'],
        },
        {
            title: 'Features',
            description: ['Cool stuff', 'Random feature', 'Team feature', 'Developer stuff', 'Another one'],
        },
        {
            title: 'Resources',
            description: ['Resource', 'Resource name', 'Another resource', 'Final resource'],
        },
        {
            title: 'Legal',
            description: ['Privacy policy', 'Terms of use'],
        },
    ];

    return (
        <footer className='footer'>
            <Grid container spacing={2} justify="space-evenly">
                {footers.map(footer => (
                    <Grid item xs key={footer.title}>
                        <Typography variant="h6" color="textPrimary" gutterBottom>
                            {footer.title}
                        </Typography>
                        {footer.description.map(item => (
                            <Typography key={item} variant="subtitle1" color="textSecondary">
                                {item}
                            </Typography>
                        ))}
                    </Grid>
                ))}
            </Grid>
        </footer>
    );
};

export default Footer;