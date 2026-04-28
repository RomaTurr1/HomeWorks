import {Card, CardContent, Typography, Button} from "@mui/material"

export default function ProductCardMui(){
    return(
        <Card sx = {{maxWidth: 320, mb: 2}}>
            <CardContent>
                <Typography variant="h6">
                    Ноутбук
                </Typography>
                <Typography variant="body1" sx = {{mb:2}}>
                    350000тг
                </Typography>
                <Button variant="container">
                    Добавить в корзину
                </Button>
            </CardContent>
        </Card>
    )
}