We will continue <i>Change Bidirectional Association to Unidirectional</i> from the place where we stopped in the inverse refactoring.

We have <code>Customer</code> and <code>Order</code> classes with a bidirectional association.

Since completion of the previous refactoring technique, two new methods have appeared in the code:

Method for getting price with discount in the order class

Recently we received new requirements: orders must appear only if the customer has already been created. This lets us forego bidirectional association between the classes and get rid of the association between the order and customer.

This works particularly well when behavior is called by client code containing a customer object that can be passed as an argument.

Another alternative is to change the method for getting data so as to get the customer while bypassing use of any field. Then you can use <a href="/substitute-algorithm">Substitute Algorithm</a> on the body of <code>Order.getCustomer</code> and do something similar to the following.

The previous introduction of a parameter into the method can now be removed, since the customer getter will return the correct object.

Slow… but it works. In the context of a database, things may even become a little faster if a database query is used.

Now prepare to remove the <code>setCustomer</code> method by replacing their calls in the code of the customer class with direct addition of order objects to the collection.

Then remove the method that assigns a new customer value in the order class.

Now just remove the field itself, fully eliminating the bidirectional association between the classes.

Let's start the final testing.

Now refactoring is complete. If you like, you can compare the old and new code.