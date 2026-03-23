package repo

import "fmt"

func (p *CreateOrderItemParams) Validate() error {
	if p.ProductID == 0 {
		return fmt.Errorf("Product ID is required")
	}
	if p.Quantity == 0 {
		return fmt.Errorf("Quantity is required")
	}
	return nil
}

func (p *CreateOrderParams) Validate() error {
	if p.CustomerID == 0 {
		return fmt.Errorf("Customer ID is required")
	}
	if p.Status == "" {
		return fmt.Errorf("Status is required")
	}
	if len(p.Items) == 0 {
		return fmt.Errorf("at least one item is required")
	}

	return nil
}

func (p *UpdateProductPriceParams) Validate() error {
	if p.ID == 0 {
		return fmt.Errorf("Product ID is required")
	}
	if p.PriceInCents == 0 {
		return fmt.Errorf("Price is required")
	}
	return nil
}

func (p *UpdateProductWhereIDParams) Validate() error {
	if p.ID == 0 {
		return fmt.Errorf("Product ID is required")
	}
	if p.Name == "" {
		return fmt.Errorf("Name is required")
	}
	if p.PriceInCents == 0 {
		return fmt.Errorf("Price is required")
	}
	if p.Quantity == 0 {
		return fmt.Errorf("Quantity is required")
	}
	return nil
}

func (p *CreateProductParams) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("Name is required")
	}
	if p.PriceInCents == 0 {
		return fmt.Errorf("Price is required")
	}
	if p.Quantity == 0 {
		return fmt.Errorf("Quantity is required")
	}
	return nil
}

func (p *IncreaseProductQuantityParams) Validate() error {
	if p.ID == 0 {
		return fmt.Errorf("Product ID is required")
	}
	if p.Quantity == 0 {
		return fmt.Errorf("Product Quantity is required")
	}
	return nil
}
