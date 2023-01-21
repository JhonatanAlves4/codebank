import Head from "next/head";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  CardMedia,
  Typography,
} from "@material-ui/core";
import { Product } from "@/model";
import { GetStaticPaths, GetStaticProps, NextPage } from "next";
import http from "@/http";

interface ProductDetailPageProps {
  product: Product;
}

const ProductDetailPage: NextPage<ProductDetailPageProps> = ({product}) => {
  return (
    <>
      <Head>
        <title>{product.name} - Detalhes do produto</title>
        <meta name="description" content="Generated by create next app" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Card>
        <CardHeader
          title={product.name.toUpperCase()}
          subheader={`R$ ${product.price}`}
        />
        <CardActions>
          <Button size="small" color="primary" component="a">
            Comprar
          </Button>
        </CardActions>
        <CardMedia style={{paddingTop: '50%'}} image={product.image_url} />
        <CardContent>
          <Typography variant="body2" color="textSecondary" component="p">
            {product.description}
          </Typography>
        </CardContent>
      </Card>
    </>
  );
};

export default ProductDetailPage;

export const getStaticProps: GetStaticProps<ProductDetailPageProps, {slug: string}> = async (context) => {
  const { slug } = context.params!;
  const { data: product } = await http.get(`products/${slug}`);
  
  return {
    props: {
      product,
      //products: response.data
    },
    revalidate: 1 * 60 * 2
  };
};

export const getStaticPaths: GetStaticPaths = async (context) => {
  const { data: products } = await http.get(`products`);

  const paths = products.map((p: Product) => ({
    params: {slug: p.slug}
  }))

  return {paths, fallback: "blocking"}
}