//
//  ProfileViewController.swift
//  Mode
//
//  Created by Jackson Tubbs on 9/6/19.
//  Copyright © 2019 Jax Tubbs. All rights reserved.
//

import UIKit

class ProfileViewController: UIViewController {

    // MARK: - Outlets
    
//    @IBOutlet weak var profileNavigationBar: UINavigationBar!
//    @IBOutlet weak var profileHandleNavigationItem: UINavigationItem!
    @IBOutlet weak var profileCollectionView: UICollectionView!
    @IBOutlet weak var profileNameLabel: UILabel!
    @IBOutlet weak var profileHandleLabel: UILabel!
    @IBOutlet weak var profileImageImageView: UIImageView!
    
    // MARK: - Properties
    
    
    // MARK: - Lifecycle
    
    override func viewDidLoad() {
        super.viewDidLoad()
        profileCollectionView.delegate = self
        profileCollectionView.dataSource = self
        updateViews()
    }
    
    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(animated)
    }
    override func viewWillAppear(_ animated: Bool) {
        super.viewWillAppear(animated)

        updateViews()
    }
    
    override func viewDidLayoutSubviews() {
        super.viewDidLayoutSubviews()
//        profileImageImageView.layer.cornerRadius = profileImageImageView.frame.height
    }
    
    // MARK: - Actions
    
    // MARK: - Custom Functions
    
    func updateHandle(handle: String) {
        profileHandleLabel.text = handle
    }
    
    func updateName(name: String) {
        profileNameLabel.text = name
    }
    
    func updateProfilePhoto(image: UIImage) {
        profileImageImageView.image = image
    }
    
    func updateViews() {
        updateHandle(handle: "@jpwashman")
        updateName(name: "Josh Porter")
        updateProfilePhoto(image: UIImage(named: "exampleProfilePhoto.jpg")!)
        
        profileImageImageView.layer.cornerRadius = profileImageImageView.frame.height / 2
        
        // Removes vertical spacing between cells
//        let layout: UICollectionViewFlowLayout = UICollectionViewFlowLayout()
//        layout.minimumLineSpacing = 0
//        layout.minimumInteritemSpacing = 0
//        profileCollectionView.collectionViewLayout = layout
    }

    /*
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        // Get the new view controller using segue.destination.
        // Pass the selected object to the new view controller.
    }
    */
} // End of class

// MARK: - Extensions

extension ProfileViewController: UICollectionViewDelegate, UICollectionViewDataSource, UICollectionViewDelegateFlowLayout {
    
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return 30
    }
    
    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        guard let cell = profileCollectionView.dequeueReusableCell(withReuseIdentifier: "photoCell", for: indexPath) as? PostCollectionViewCell else {return UICollectionViewCell()}
        
        var imageName: String = ""
        
        let imageNumber = Int.random(in: 0..<5)
        
        switch imageNumber {
        case 0: imageName = "profileOne"
        case 1: imageName = "profileTwo"
        case 2: imageName = "profileThree"
        case 3: imageName = "profileFour"
        case 4: imageName = "profileFive"
        default: fatalError("Big Error at \(#function)")
        }
        
        let image = UIImage(named: imageName)!
        cell.image = image
        
        return cell
    }
    
    // MARK: - FlowLayout
    
    func collectionView(_ collectionView: UICollectionView, layout collectionViewLayout: UICollectionViewLayout, sizeForItemAt indexPath: IndexPath) -> CGSize {
        let size = CGSize(width: profileCollectionView.frame.width / 3 - 1, height: profileCollectionView.frame.width / 3 - 1)
        return size
    }
}
