from re import T


def wrap(func):

    def inner(obj):
        print("Wrapped as a charm")
        func(obj)
    
    return inner
    
    
class Animals:
    
    def __init__(self):
        self.isAnimal = True

    def __str__(self):
        return "Parent class for all animals"
    
    def animal(self):
        print('I am animal')
    
    
class Bird(Animals):
    
    def __init__(self):
        super().__init__()
        self.speed = 250
        self.canFly = True
        
    @wrap
    def animal(self):
        print('I am Bird')    
    
    def __str__(self):
        return "Dedicated bird class"
    
# b1 = Bird()
# b1.animal()


def returnTrue() -> bool:
    print("Return True")
    return True

    
    
def zbabir():
    try:
        print("Inside the try")
        input("Press to close\n")
        fadlfkajsdf
    except:
        zbabir()
    
zbabir()