import React from "react";
import {useState} from "react";
import {useDispatch} from "react-redux";
import {useNavigate} from "react-router-dom";
import {Box, Typography, Button, Stack, ButtonGroup, List, ListItem, ListItemText, Divider, drawerClasses} from "@mui/material"
import { createTheme, ThemeProvider } from "@mui/material/styles";
import {DeleteForeverOutlined, Add, Remove} from "@mui/icons-material"

import { removeCartItem, removeAllCartItems, addCartItem, checkout } from "../../../actions";
import { formatCurrency } from "../../../helpers/Utils";

const FilledCart = ({cart, drawer}) => {
    const dispatch = useDispatch()
    const navigate = useNavigate()

    const disabledButtonTheme = createTheme({
        palette: {
            action: {
                disabledBackground: 'black',
                disabled: 'inherit',
            }
        }
    });

    const theme = createTheme({
        palette: {
            primary: {
                main: "#212121",
            },
        },
    });

    const handleQty = (op, productId, qty) => {
        // Add +1 to qty if item already exists in cart
        let obj = cart.products.find(item => item.product_id === productId)

        if (obj.qty !== undefined){
            if (op === "increment"){
                qty = obj.qty + qty
            } else if (op === "decrement"){
                qty = obj.qty - qty
            }
        }

        if (qty === 0) {
            dispatch(removeCartItem(productId))
        } else {
            dispatch(addCartItem(productId, qty))
        }
    }

    const handleCheckout = () => {
        dispatch(checkout()).then(() => {
            drawer.onClick(false)
            navigate("/payment")
        })
    }

    return (
        <Stack directon="row" spacing={2}>
            <List sx={{
                maxHeight: "30vw",
                overflow: "auto",
            }}>
                {cart.products.map((product) => (
                    <ListItem key={product.product_id}>
                        <Box component="img" sx={{
                                    minWidth: {xs: 50, md: 75},
                                    maxHeight: { xs: 50, md: 75 },
                                    maxWidth: { xs: 50, md: 75 },
                                    }}
                                    alt="product img"
                                    src={product.image}/>
                        <ListItemText sx={{
                            pl: 2,
                            pr: 2,
                            overflowWrap: "break-word",
                        }}
                        primary={product.name}
                        secondary={`$${formatCurrency(product.price)}`}/>
                        <div>
                            <Typography variant="subtitle1" align="right">
                                ${product.qty === undefined ? 0 : formatCurrency(product.price * product.qty)}
                            </Typography>
                            <ButtonGroup variant="outlined" size="small" aria-label="outlined primary button group">
                                <Button color="inherit" onClick={() => handleQty(
                                    "increment",
                                    product.product_id,
                                    1,
                                )}><Add/></Button>
                                <Button color="inherit" onClick={() => handleQty(
                                    "decrement",
                                    product.product_id,
                                    1,
                                )}><Remove/></Button>
                                <ThemeProvider theme={disabledButtonTheme}>
                                    <Button disabled>{product.qty === undefined ? 0 : product.qty}</Button>
                                </ThemeProvider>
                                <Button color="inherit" onClick={() => dispatch(removeCartItem(product.product_id))}>
                                    <DeleteForeverOutlined/>
                                </Button>
                                
                            </ButtonGroup>
                        </div>
                    </ListItem>
                ))}
                <Divider/>
            </List>
            <Box sx={{
                display:"flex",
                justifyContent: "space-between",
            }}>
                <Typography variant="subtitle1">Subtotal</Typography>
                <Typography variant="subtitle1">${formatCurrency(cart.subtotal)}</Typography>
            </Box>
            <ThemeProvider theme={theme}>
                <Button type="button" variant="contained" color="primary" fullWidth onClick={handleCheckout}>Checkout</Button>
            </ThemeProvider>
            <Button type="button" variant="outlined" color="inherit" fullWidth onClick={() => dispatch(removeAllCartItems())}>Empty Cart</Button>
        </Stack>
    )
}

export default FilledCart