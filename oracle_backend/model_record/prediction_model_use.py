import numpy as np
import torch


def prediction(data_one):
    tree = torch.load("model01.pth")
    data_two = []
    data_two.append(data_one)
    data_np = np.array(data_two)
    predicton = tree.predict(data_np)
    result = predicton.tolist()
    return result[0]


