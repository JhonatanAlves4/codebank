import Head from "next/head";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  CardMedia,
  Grid,
  Typography,
} from "@material-ui/core";
import { Product, products } from "@/model";
import Link from "next/link";

export default function ProductsListPage() {
  return (
    <>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Typography component="h1" variant="h3" color="textPrimary" gutterBottom>
        Produtos
      </Typography>
      <Grid container spacing={4}>
        {products.map((product) => (
          <Grid key={product.id} item xs={12} sm={6} md={4}>
            <Card>
              <CardMedia
                style={{ paddingTop: "50%" }}
                image={product.image_url}
              />
              <CardContent>
                <Typography component="h2" variant="h5" gutterBottom>
                  {product.name}
                </Typography>
              </CardContent>
              <CardActions>
                <Link
                  href="/products/[slug]"
                  as={`/products/${product.slug}`}
                  passHref
                >
                  <Button size="small" color="primary" component="a">
                    Detalhes
                  </Button>
                </Link>
              </CardActions>
            </Card>
          </Grid>
        ))}
      </Grid>
    </>
  );
}