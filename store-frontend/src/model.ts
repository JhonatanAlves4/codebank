export interface Product {
  id: string;
  name: string;
  description: string;
  image_url: string;
  price: number;
  slug: string;
  created_at: string;
}

export const products: Product[] = [
  {
    id: "uuid",
    name: "produto",
    description: "descrição produto",
    price: 49.99,
    image_url: "https://source.unsplash.com/random?product",
    slug: "produto-teste",
    created_at: "2023-01-19T00:00:00",
  },
];