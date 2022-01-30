import React from "react";
import {Toolbar, IconButton, Badge, MenuItem, Menu, Typography, AppBar, TextField, InputAdornment} from "@mui/material";
import {ShoppingCart, Search} from "@mui/icons-material";
import { useDispatch } from "react-redux";

import logo from "../../assets/logo.svg";
import { searchProducts } from "../../actions/Products";

const Navbar = ({totalItems, onClickDrawer}) => {
    const dispatch = useDispatch();

    const search = (e) => {
        dispatch(searchProducts(e.target.value))
    }

    return (
        <AppBar position="sticky" color="inherit">
            <Toolbar>
                <Typography variant="h6" color="inherit">
                    <img src={logo} alt="Shopping-Micro" height="25px" />
                </Typography>
                <div>
                    <TextField
                        id="outlined-adornment-password"
                        InputProps={{
                            startAdornment: (
                                <InputAdornment position="start">
                                    <Search />
                                </InputAdornment>
                            ),
                        }}
                        placeholder="Search..."
                        onChange={search}
                        size="small"
                        variant="outlined"
                    />
                </div>
                <div>
                    <IconButton aria-label="Show cart items" color="inherit" onClick={() => onClickDrawer(true)}>
                        <Badge badgeContent={totalItems} color="secondary">
                            <ShoppingCart />
                        </Badge>
                    </IconButton>
                </div>
            </Toolbar>
        </AppBar>
    )
}

export default React.memo(Navbar)