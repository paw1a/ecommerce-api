import * as React from 'react';
import image from './img.png'
import BookmarkIcon from '@mui/icons-material/Bookmark';
import './Card.css'
import IconButton from "@mui/material/IconButton";
import StarIcon from '@mui/icons-material/Star';
import ModeCommentIcon from '@mui/icons-material/ModeComment';
import AddShoppingCartIcon from '@mui/icons-material/AddShoppingCart';

export default function ProductCard({product}) {
    return (
        <div className='product-card'>
            <div className='image-row'>
                <img src={image} alt="image"/>
                <IconButton style={{paddingLeft: 0}} disableRipple='true'>
                    <BookmarkIcon className='card-button'/>
                </IconButton>

                <div>
                    <IconButton disableRipple='true' style={{fontSize: '14px'}}>
                        <StarIcon className='card-button'/>
                        4.5
                    </IconButton>

                    <IconButton style={{fontSize: '13px'}} disableRipple='true'>
                        <ModeCommentIcon className='card-button' fontSize='small'/>
                        17
                    </IconButton>
                </div>
            </div>

            <h4>{product.name}</h4>
            <div className='price-cart'>
                <h2>{product.price} $</h2>
                <IconButton disableRipple='true'>
                    <AddShoppingCartIcon className='card-button' fontSize='large'/>
                </IconButton>
            </div>
        </div>
    );
}
