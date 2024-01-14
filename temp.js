const express = require('express')

const app = express()

app.get('/list-products', checkAuth, getProducts)


const getProducts = (req, res) => {
    res.status(200).json({
        products: [
            /*
                Products Data
            */
        ],
    })
}

const checkAuth = (req, res, next) => {
    try {
        if(!req.headers.Bearer) {
            res.status(401).json({
                error : 'You are not logged in.'
            })
            return
        }
        
        // Some logic to decode token
        const user = decodeToken(req.headers.Bearer)
    
        req.loggedInUser = user
        next()
    } catch(err) {
        res.status(500).json({
            error: 'Internal Server Error'
        })
    }
}

const decodeToken = (token) => {
    // JWT Token usko decoded krke user ka account return kr dega

    // Verification Done
    return {
        name : 'Sushant',
        email : 'sushant@sushant.com'
    }
}

app.listen(3000, () => {
    console.log('Listening')
})