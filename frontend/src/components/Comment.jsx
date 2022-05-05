import * as React from 'react';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import CardContent from '@mui/material/CardContent';
import Avatar from '@mui/material/Avatar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import StarIcon from '@mui/icons-material/Star';

import './Comment.css'

export default function Comment({comment}) {
    return (
        <Card sx={{backgroundColor: '#f5f5f5', boxShadow: 'none'}}>
            <CardHeader
                avatar={
                    <Avatar aria-label="recipe"/>
                }
                title={comment.username}
                subheader={comment.date}
                action={
                    <IconButton aria-label="settings" disabled>
                        <StarIcon/>
                        {comment.rating}
                    </IconButton>
                }
                sx={{paddingLeft: 0}}
            />
            <CardContent sx={{paddingLeft: 0}}>
                <Typography variant="body2" color="text.secondary">
                    {comment.text}
                </Typography>
            </CardContent>
        </Card>
    );
}
