import * as React from 'react';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardContent from '@mui/material/CardContent';
import Avatar from '@mui/material/Avatar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import StarIcon from '@mui/icons-material/Star';

import './Comment.css'

export default function Comment() {
    return (
        <Card sx={{backgroundColor: '#f5f5f5', boxShadow: 'none'}}>
            <CardHeader
                avatar={
                    <Avatar aria-label="recipe"/>
                }
                title="Shrimp and Chorizo Paella"
                subheader="September 14, 2016"
                action={
                    <IconButton aria-label="settings" disabled>
                        <StarIcon/>
                        5
                    </IconButton>
                }
                sx={{paddingLeft: 0}}
            />
            <CardContent sx={{paddingLeft: 0}}>
                <Typography variant="body2" color="text.secondary">
                    This impressive paella is a perfect party dish and a fun meal to cook
                    together with your guests. Add 1 cup of frozen peas along with the mussels,
                    if you like.
                </Typography>
            </CardContent>
        </Card>
    );
}
