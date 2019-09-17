//
//  CatagoryCollectionViewCell.swift
//  Mode
//
//  Created by Jackson Tubbs on 9/16/19.
//  Copyright © 2019 Jax Tubbs. All rights reserved.
//

import UIKit

class CatagoryCollectionViewCell: UICollectionViewCell {
    
    // MARK: - Outlets
    
    @IBOutlet weak var buttonBackgroundView: UIView!
    @IBOutlet weak var catagoryButton: UIButton!
    
    // MARK: - Lifecycle
    
    override func awakeFromNib() {
        updateViews()
    }
    
    // MARK: - Actions
    
    @IBAction func catagoryButtonTapped(_ sender: Any) {
        print("Button Tapped")
    }
    
    // MARK: - Custom Functions
    
    func updateViews() {
        buttonBackgroundView.layer.cornerRadius = 35 / 2
        buttonBackgroundView.layer.borderColor = UIColor.black.cgColor
        buttonBackgroundView.layer.borderWidth = 1
    }
}
