import * as React from 'react';
import image from './img.png'
import BookmarkIcon from '@mui/icons-material/Bookmark';
import './Card.css'
import IconButton from "@mui/material/IconButton";
import StarIcon from '@mui/icons-material/Star';
import ModeCommentIcon from '@mui/icons-material/ModeComment';
import AddShoppingCartIcon from '@mui/icons-material/AddShoppingCart';
import Link from "@mui/material/Link";

export default function ProductCard({product}) {
    return (
        <div className='product-card'>
            <div className='image-row'>
                <Link href={"/product/" + product.id}>
                    <img src={image} alt="image"/>
                </Link>
                <IconButton style={{paddingLeft: 0}} disableRipple>
                    <BookmarkIcon className='card-button'/>
                </IconButton>

                <div>
                    <IconButton disableRipple style={{fontSize: '14px'}}>
                        <StarIcon className='card-button'/>
                        {product.totalRating}
                    </IconButton>

                    <IconButton style={{fontSize: '13px'}} disableRipple>
                        <ModeCommentIcon className='card-button' fontSize='small'/>
                        17
                    </IconButton>
                </div>
            </div>

            <Link color='inherit' href="/product/1" style={{textDecoration: 'none'}}>
                <h4>{product.name}</h4>
            </Link>
            <div className='price-cart'>
                <h2>{product.price} $</h2>
                <IconButton disableRipple>
                    <AddShoppingCartIcon className='card-button' fontSize='large'/>
                </IconButton>
            </div>
        </div>
    );
}
