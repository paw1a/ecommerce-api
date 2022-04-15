import React from 'react';
import {Button, Card, CardActions, CardContent, Typography} from "@mui/material";

const ProductItem = ({product}) => {
    return (
        <Card variant="outlined" sx={{ minWidth: 275 }}>
            <CardContent>
                <Typography sx={{ fontSize: 24 }} color="text.primary" gutterBottom>
                    {product.name}
                </Typography>
                <Typography variant="h6" component="div">
                    {product.price} $
                </Typography>
                <Typography sx={{ mb: 1.5 }} color="text.secondary">
                    Rating: {product.totalRating}
                </Typography>
                <Typography variant="body2">
                    {product.description}
                </Typography>
            </CardContent>
            <CardActions>
                <Button size="small">Learn More</Button>
            </CardActions>
        </Card>
    );
};

export default ProductItem;